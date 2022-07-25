package ipc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/kom0055/go-hadoop/common/defined"
	"github.com/kom0055/go-hadoop/common/security"
	"github.com/kom0055/go-hadoop/proto/common"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
	"log"
	"net"
	"strings"
	"sync"
)

type Client struct {
	ClientId      *uuid.UUID
	Ugi           *common.UserInformationProto
	ServerAddress string
	TCPNoDelay    bool
}

type connection struct {
	con *net.TCPConn
}

type connIdentity struct {
	user     string
	protocol string
	address  string
}

type call struct {
	callId     int32
	procedure  proto.Message
	request    proto.Message
	response   proto.Message
	err        *error
	retryCount int32
}

func (c *Client) String() string {
	buf := bytes.NewBufferString("")
	_, _ = fmt.Fprint(buf, "<clientId:", c.ClientId)
	_, _ = fmt.Fprint(buf, ", server:", c.ServerAddress)
	_, _ = fmt.Fprint(buf, ">")
	return buf.String()
}

var (
	SaslRpcDummyClientId           = make([]byte, 0)
	SaslRpcCallId            int32 = -33
	SaslRpcInvalidRetryCount int32 = -1
)

func (c *Client) Call(ctx context.Context, rpc *common.RequestHeaderProto, rpcRequest proto.Message, rpcResponse proto.Message) error {
	// Create connIdentity
	connectionId := connIdentity{user: *c.Ugi.RealUser, protocol: *rpc.DeclaringClassProtocolName, address: c.ServerAddress}

	// Get connection to server
	//log.Println("Connecting...", c)
	conn, err := getConnection(c, &connectionId)
	if err != nil {
		return err
	}

	// Create call and send request
	rpcCall := call{callId: 0, procedure: rpc, request: rpcRequest, response: rpcResponse}
	err = sendRequest(c, conn, &rpcCall)
	if err != nil {
		log.Println("sendRequest", err)
		return err
	}

	// Read & return response
	err = c.readResponse(conn, &rpcCall)

	return err
}

var connectionPool = struct {
	sync.RWMutex
	connections map[connIdentity]*connection
}{connections: make(map[connIdentity]*connection)}

func findUsableTokenForService(service string) (*common.TokenProto, bool) {
	userTokens := security.GetCurrentUser().GetUserTokens()

	log.Printf("looking for token for service: %s\n", service)

	if len(userTokens) == 0 {
		return nil, false
	}

	token := userTokens[service]
	if token != nil {
		return token, true
	}

	return nil, false
}

func getConnection(c *Client, connectionId *connIdentity) (*connection, error) {
	// Try to re-use an existing connection
	connectionPool.RLock()
	con := connectionPool.connections[*connectionId]
	connectionPool.RUnlock()

	// If necessary, create a new connection and save it in the connection-pool
	var err error
	if con == nil {
		con, err = setupConnection(c)
		if err != nil {
			log.Println("Couldn't setup connection: ", err)
			return nil, err
		}

		connectionPool.Lock()
		connectionPool.connections[*connectionId] = con
		connectionPool.Unlock()

		authProtocol := defined.AUTH_PROTOCOL_NONE

		if _, found := findUsableTokenForService(c.ServerAddress); found {
			log.Printf("found token for service: %s", c.ServerAddress)
			authProtocol = defined.AUTH_PROTOCOL_SASL
		}

		_ = writeConnectionHeader(con, authProtocol)

		if authProtocol == defined.AUTH_PROTOCOL_SASL {
			log.Println("attempting SASL negotiation.")

			if err = negotiateSimpleTokenAuth(c, con); err != nil {
				log.Println("failed to complete SASL negotiation!")
				return nil, err
			}

		} else {
			log.Println("no usable tokens. proceeding without auth.")
		}

		_ = writeConnectionContext(c, con, connectionId, authProtocol)
	}

	return con, nil
}

func setupConnection(c *Client) (*connection, error) {
	addr, _ := net.ResolveTCPAddr("tcp", c.ServerAddress)
	tcpConn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("error: ", err)
		return nil, err
	} else {
		log.Println("Successfully connected ", c)
	}

	// TODO: Ping thread

	// Set tcp no-delay
	_ = tcpConn.SetNoDelay(c.TCPNoDelay)

	return &connection{tcpConn}, nil
}

func writeConnectionHeader(conn *connection, authProtocol defined.AuthProtocol) error {
	// RPC_HEADER
	if _, err := conn.con.Write(defined.RpcHeader); err != nil {
		log.Println("conn.Write gohadoop.RPC_HEADER", err)
		return err
	}

	// RPC_VERSION
	if _, err := conn.con.Write(defined.Version); err != nil {
		log.Println("conn.Write gohadoop.VERSION", err)
		return err
	}

	// RPC_SERVICE_CLASS
	if serviceClass, err := defined.ConvertFixedToBytes(defined.RpcServiceClass); err != nil {
		log.Println("binary.Write", err)
		return err
	} else if _, err := conn.con.Write(serviceClass); err != nil {
		log.Println("conn.Write RPC_SERVICE_CLASS", err)
		return err
	}

	// AuthProtocol
	if authProtocolBytes, err := defined.ConvertFixedToBytes(authProtocol); err != nil {
		log.Println("WTF AUTH_PROTOCOL", err)
		return err
	} else if _, err := conn.con.Write(authProtocolBytes); err != nil {
		log.Println("conn.Write gohadoop.AUTH_PROTOCOL", err)
		return err
	}

	return nil
}

func writeConnectionContext(c *Client, conn *connection, connectionId *connIdentity, _ defined.AuthProtocol) error {
	// Create hadoop_common.IpcConnectionContextProto
	ugi, _ := defined.CreateSimpleUGIProto()
	ipcCtxProto := common.IpcConnectionContextProto{UserInfo: ugi, Protocol: &connectionId.protocol}

	// Create RpcRequestHeaderProto
	callId := int32(-3)
	clientId := [16]byte(*c.ClientId)

	/*if (authProtocol == gohadoop.AUTH_PROTOCOL_SASL) {
	  callId = SASL_RPC_CALL_ID
	}*/

	rpcReqHeaderProto := common.RpcRequestHeaderProto{RpcKind: &defined.RpcProtocolBufffer, RpcOp: &defined.RpcFinalPacket, CallId: &callId, ClientId: clientId[0:16], RetryCount: &defined.RpcDefaultRetryCount}

	rpcReqHeaderProtoBytes, err := proto.Marshal(&rpcReqHeaderProto)
	if err != nil {
		log.Println("proto.Marshal(&rpcReqHeaderProto)", err)
		return err
	}

	ipcCtxProtoBytes, _ := proto.Marshal(&ipcCtxProto)
	if err != nil {
		log.Println("proto.Marshal(&ipcCtxProto)", err)
		return err
	}

	totalLength := len(rpcReqHeaderProtoBytes) + sizeVarint(len(rpcReqHeaderProtoBytes)) + len(ipcCtxProtoBytes) + sizeVarint(len(ipcCtxProtoBytes))
	tLen := int32(totalLength)
	totalLengthBytes, err := defined.ConvertFixedToBytes(tLen)

	if err != nil {
		log.Println("ConvertFixedToBytes(totalLength)", err)
		return err
	} else if _, err := conn.con.Write(totalLengthBytes); err != nil {
		log.Println("conn.con.Write(totalLengthBytes)", err)
		return err
	}

	if err := writeDelimitedBytes(conn, rpcReqHeaderProtoBytes); err != nil {
		log.Println("writeDelimitedBytes(conn, rpcReqHeaderProtoBytes)", err)
		return err
	}
	if err := writeDelimitedBytes(conn, ipcCtxProtoBytes); err != nil {
		log.Println("writeDelimitedBytes(conn, ipcCtxProtoBytes)", err)
		return err
	}

	return nil
}

func sizeVarint(x int) (n int) {
	for {
		n++
		x >>= 7
		if x == 0 {
			break
		}
	}
	return n
}

func sendRequest(c *Client, conn *connection, rpcCall *call) error {
	//log.Println("About to call RPC: ", rpcCall.procedure)

	// 0. RpcRequestHeaderProto
	clientId := *c.ClientId
	rpcReqHeaderProto := common.RpcRequestHeaderProto{RpcKind: &defined.RpcProtocolBufffer, RpcOp: &defined.RpcFinalPacket, CallId: &rpcCall.callId, ClientId: clientId[0:16], RetryCount: &rpcCall.retryCount}
	rpcReqHeaderProtoBytes, err := proto.Marshal(&rpcReqHeaderProto)
	if err != nil {
		log.Println("proto.Marshal(&rpcReqHeaderProto)", err)
		return err
	}

	// 1. RequestHeaderProto
	requestHeaderProto := rpcCall.procedure
	requestHeaderProtoBytes, err := proto.Marshal(requestHeaderProto)
	if err != nil {
		log.Println("proto.Marshal(&requestHeaderProto)", err)
		return err
	}

	// 2. Param
	paramProto := rpcCall.request
	paramProtoBytes, err := proto.Marshal(paramProto)
	if err != nil {
		log.Println("proto.Marshal(&paramProto)", err)
		return err
	}

	totalLength := len(rpcReqHeaderProtoBytes) + sizeVarint(len(rpcReqHeaderProtoBytes)) + len(requestHeaderProtoBytes) + sizeVarint(len(requestHeaderProtoBytes)) + len(paramProtoBytes) + sizeVarint(len(paramProtoBytes))
	tLen := int32(totalLength)
	if totalLengthBytes, err := defined.ConvertFixedToBytes(tLen); err != nil {
		log.Println("ConvertFixedToBytes(totalLength)", err)
		return err
	} else {
		if _, err := conn.con.Write(totalLengthBytes); err != nil {
			log.Println("conn.con.Write(totalLengthBytes)", err)
			return err
		}
	}

	if err := writeDelimitedBytes(conn, rpcReqHeaderProtoBytes); err != nil {
		log.Println("writeDelimitedBytes(conn, rpcReqHeaderProtoBytes)", err)
		return err
	}
	if err := writeDelimitedBytes(conn, requestHeaderProtoBytes); err != nil {
		log.Println("writeDelimitedBytes(conn, requestHeaderProtoBytes)", err)
		return err
	}
	if err := writeDelimitedBytes(conn, paramProtoBytes); err != nil {
		log.Println("writeDelimitedBytes(conn, paramProtoBytes)", err)
		return err
	}

	//log.Println("Succesfully sent request of length: ", totalLength)

	return nil
}

func writeDelimitedTo(conn *connection, msg proto.Message) error {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		log.Println("proto.Marshal(msg)", err)
		return err
	}
	return writeDelimitedBytes(conn, msgBytes)
}

func writeDelimitedBytes(conn *connection, data []byte) error {
	if _, err := conn.con.Write(protowire.AppendVarint([]byte{}, uint64(len(data)))); err != nil {
		log.Println("conn.con.Write(protowire.AppendVarint(uint64(len(data))))", err)
		return err
	}

	//if _, err := conn.con.Write(proto.EncodeVarint(uint64(len(data)))); err != nil {
	//	log.Fatal("conn.con.Write(proto.EncodeVarint(uint64(len(data))))", err)
	//	return err
	//}
	if _, err := conn.con.Write(data); err != nil {
		log.Println("conn.con.Write(data)", err)
		return err
	}

	return nil
}

func (c *Client) readResponse(conn *connection, rpcCall *call) error {
	// Read first 4 bytes to get total-length
	var totalLength int32 = -1
	var totalLengthBytes [4]byte
	if _, err := conn.con.Read(totalLengthBytes[0:4]); err != nil {
		log.Println("conn.con.Read(totalLengthBytes)", err)
		return err
	}

	if err := defined.ConvertBytesToFixed(totalLengthBytes[0:4], &totalLength); err != nil {
		log.Println("gohadoop.ConvertBytesToFixed(totalLengthBytes, &totalLength)", err)
		return err
	}

	responseBytes := make([]byte, totalLength)
	if _, err := conn.con.Read(responseBytes); err != nil {
		log.Println("conn.con.Read(totalLengthBytes)", err)
		return err
	}

	// Parse RpcResponseHeaderProto
	rpcResponseHeaderProto := common.RpcResponseHeaderProto{}
	off, err := readDelimited(responseBytes[0:totalLength], &rpcResponseHeaderProto)
	if err != nil {
		log.Println("readDelimited(responseBytes, rpcResponseHeaderProto)", err)
		return err
	}
	//log.Println("Received rpcResponseHeaderProto = ", rpcResponseHeaderProto)

	err = c.checkRpcHeader(&rpcResponseHeaderProto)
	if err != nil {
		log.Println("c.checkRpcHeader failed", err)
		return err
	}

	if *rpcResponseHeaderProto.Status == common.RpcResponseHeaderProto_SUCCESS {
		// Parse RpcResponseWrapper
		_, err = readDelimited(responseBytes[off:], rpcCall.response)
	} else {
		log.Println("RPC failed with status: ", rpcResponseHeaderProto.Status.String())
		errorDetails := [4]string{rpcResponseHeaderProto.Status.String(), "ServerDidNotSetExceptionClassName", "ServerDidNotSetErrorMsg", "ServerDidNotSetErrorDetail"}
		if rpcResponseHeaderProto.ExceptionClassName != nil {
			errorDetails[0] = *rpcResponseHeaderProto.ExceptionClassName
		}
		if rpcResponseHeaderProto.ErrorMsg != nil {
			errorDetails[1] = *rpcResponseHeaderProto.ErrorMsg
		}
		if rpcResponseHeaderProto.ErrorDetail != nil {
			errorDetails[2] = rpcResponseHeaderProto.ErrorDetail.String()
		}
		err = errors.New(strings.Join(errorDetails[:], ":"))
	}
	return err
}

func readDelimited(rawData []byte, msg proto.Message) (int, error) {
	headerLength, off := protowire.ConsumeVarint(rawData)
	if off == 0 {
		log.Println("protowire.ConsumeVarint(rawData) returned zero")
		return -1, nil
	}

	//headerLength, off := proto.DecodeVarint(rawData)
	//if off == 0 {
	//	log.Fatal("proto.DecodeVarint(rawData) returned zero")
	//	return -1, nil
	//}

	err := proto.Unmarshal(rawData[off:off+int(headerLength)], msg)
	if err != nil {
		log.Println("proto.Unmarshal(rawData[off:off+headerLength]) ", err)
		return -1, err
	}

	return off + int(headerLength), nil
}

func (c *Client) checkRpcHeader(rpcResponseHeaderProto *common.RpcResponseHeaderProto) error {
	callClientId := *c.ClientId
	headerClientId := rpcResponseHeaderProto.ClientId
	if len(rpcResponseHeaderProto.ClientId) > 0 {
		if !bytes.Equal(callClientId[0:16], headerClientId[0:16]) {
			log.Println("Incorrect clientId: ", headerClientId)
			return errors.New("incorrect clientId")
		}
	}
	return nil
}

func sendSaslMessage(c *Client, conn *connection, message *common.RpcSaslProto) error {
	saslRpcHeaderProto := common.RpcRequestHeaderProto{RpcKind: &defined.RpcProtocolBufffer,
		RpcOp:      &defined.RpcFinalPacket,
		CallId:     &SaslRpcCallId,
		ClientId:   SaslRpcDummyClientId,
		RetryCount: &SaslRpcInvalidRetryCount}

	saslRpcHeaderProtoBytes, err := proto.Marshal(&saslRpcHeaderProto)

	if err != nil {
		log.Println("proto.Marshal(&saslRpcHeaderProto)", err)
		return err
	}

	saslRpcMessageProtoBytes, err := proto.Marshal(message)

	if err != nil {
		log.Println("proto.Marshal(saslMessage)", err)
		return err
	}

	totalLength := len(saslRpcHeaderProtoBytes) + sizeVarint(len(saslRpcHeaderProtoBytes)) + len(saslRpcMessageProtoBytes) + sizeVarint(len(saslRpcMessageProtoBytes))
	tLen := int32(totalLength)

	if totalLengthBytes, err := defined.ConvertFixedToBytes(tLen); err != nil {
		log.Println("ConvertFixedToBytes(totalLength)", err)
		return err
	} else {
		if _, err := conn.con.Write(totalLengthBytes); err != nil {
			log.Println("conn.con.Write(totalLengthBytes)", err)
			return err
		}
	}
	if err := writeDelimitedBytes(conn, saslRpcHeaderProtoBytes); err != nil {
		log.Println("writeDelimitedBytes(conn, saslRpcHeaderProtoBytes)", err)
		return err
	}
	if err := writeDelimitedBytes(conn, saslRpcMessageProtoBytes); err != nil {
		log.Println("writeDelimitedBytes(conn, saslRpcMessageProtoBytes)", err)
		return err
	}

	return nil
}

func receiveSaslMessage(_ *Client, conn *connection) (*common.RpcSaslProto, error) {
	// Read first 4 bytes to get total-length
	var totalLength int32 = -1
	var totalLengthBytes [4]byte

	if _, err := conn.con.Read(totalLengthBytes[0:4]); err != nil {
		log.Println("conn.con.Read(totalLengthBytes)", err)
		return nil, err
	}
	if err := defined.ConvertBytesToFixed(totalLengthBytes[0:4], &totalLength); err != nil {
		log.Println("gohadoop.ConvertBytesToFixed(totalLengthBytes, &totalLength)", err)
		return nil, err
	}

	responseBytes := make([]byte, totalLength)

	if _, err := conn.con.Read(responseBytes); err != nil {
		log.Println("conn.con.Read(totalLengthBytes)", err)
		return nil, err
	}

	// Parse RpcResponseHeaderProto
	rpcResponseHeaderProto := common.RpcResponseHeaderProto{}
	off, err := readDelimited(responseBytes[0:totalLength], &rpcResponseHeaderProto)
	if err != nil {
		log.Println("readDelimited(responseBytes, rpcResponseHeaderProto)", err)
		return nil, err
	}

	err = checkSaslRpcHeader(&rpcResponseHeaderProto)
	if err != nil {
		log.Println("checkSaslRpcHeader failed", err)
		return nil, err
	}

	var saslRpcMessage common.RpcSaslProto

	if *rpcResponseHeaderProto.Status == common.RpcResponseHeaderProto_SUCCESS {
		// Parse RpcResponseWrapper
		if _, err = readDelimited(responseBytes[off:], &saslRpcMessage); err != nil {
			log.Println("failed to read sasl response!")
			return nil, err
		} else {
			return &saslRpcMessage, nil
		}
	} else {
		log.Println("RPC failed with status: ", rpcResponseHeaderProto.Status.String())
		errorDetails := [4]string{rpcResponseHeaderProto.Status.String(), "ServerDidNotSetExceptionClassName", "ServerDidNotSetErrorMsg", "ServerDidNotSetErrorDetail"}
		if rpcResponseHeaderProto.ExceptionClassName != nil {
			errorDetails[0] = *rpcResponseHeaderProto.ExceptionClassName
		}
		if rpcResponseHeaderProto.ErrorMsg != nil {
			errorDetails[1] = *rpcResponseHeaderProto.ErrorMsg
		}
		if rpcResponseHeaderProto.ErrorDetail != nil {
			errorDetails[2] = rpcResponseHeaderProto.ErrorDetail.String()
		}
		err = errors.New(strings.Join(errorDetails[:], ":"))
		return nil, err
	}
}

func checkSaslRpcHeader(rpcResponseHeaderProto *common.RpcResponseHeaderProto) error {
	headerClientId := rpcResponseHeaderProto.ClientId
	if rpcResponseHeaderProto.ClientId != nil {
		if !bytes.Equal(SaslRpcDummyClientId, headerClientId) {
			log.Println("Incorrect clientId: ", headerClientId)
			return errors.New("incorrect clientId")
		}
	}
	return nil
}

func negotiateSimpleTokenAuth(client *Client, con *connection) error {
	var saslNegotiateState common.RpcSaslProto_SaslState = common.RpcSaslProto_NEGOTIATE
	var saslNegotiateMessage common.RpcSaslProto = common.RpcSaslProto{State: &saslNegotiateState}
	var saslResponseMessage *common.RpcSaslProto
	var err error

	//send a SASL negotiation request
	if err = sendSaslMessage(client, con, &saslNegotiateMessage); err != nil {
		log.Println("failed to send SASL NEGOTIATE message!")
		return err
	}

	//get a response with supported mehcanisms/challenge
	if saslResponseMessage, err = receiveSaslMessage(client, con); err != nil {
		log.Println("failed to receive SASL NEGOTIATE response!")
		return err
	}

	auths := saslResponseMessage.GetAuths()

	if numAuths := len(auths); numAuths <= 0 {
		log.Println("No supported auth mechanisms!")
		return errors.New("no supported auth mechanisms")
	}

	//for now, we only support auth when TOKEN/DIGEST-MD5 is the first/only
	//supported auth mechanism
	auth := auths[0]

	if !(auth.GetMethod() == "TOKEN" && auth.GetMechanism() == "DIGEST-MD5") {
		log.Println("gohadoop only supports TOKEN/DIGEST-MD5 auth!")
		return errors.New("gohadoop only supports TOKEN/DIGEST-MD5 auth")
	}

	method := auth.GetMethod()
	mechanism := auth.GetMechanism()
	protocol := auth.GetProtocol()
	serverId := auth.GetServerId()
	challenge := auth.GetChallenge()

	//TODO: token/service mapping + token selection based on type/service
	//we wouldn't have gotten this far if there wasn't at least one available token.
	userToken, _ := findUsableTokenForService(client.ServerAddress)
	response, err := security.GetDigestMD5ChallengeResponse(protocol, serverId, challenge, userToken)

	if err != nil {
		log.Println("failed to get challenge response! ", err)
		return err
	}

	saslInitiateState := common.RpcSaslProto_INITIATE
	authSend := common.RpcSaslProto_SaslAuth{Method: &method, Mechanism: &mechanism,
		Protocol: &protocol, ServerId: &serverId}
	authsSendArray := []*common.RpcSaslProto_SaslAuth{&authSend}
	saslInitiateMessage := common.RpcSaslProto{State: &saslInitiateState,
		Token: []byte(response), Auths: authsSendArray}

	//send a SASL inititate request
	if err = sendSaslMessage(client, con, &saslInitiateMessage); err != nil {
		log.Println("failed to send SASL INITIATE message!")
		return err
	}

	//get a response with supported mehcanisms/challenge
	if saslResponseMessage, err = receiveSaslMessage(client, con); err != nil {
		log.Println("failed to read response to SASL INITIATE response!")
		return err
	}

	if saslResponseMessage.GetState() != common.RpcSaslProto_SUCCESS {
		log.Println("expected SASL SUCCESS response!")
		return errors.New("expected SASL SUCCESS response")
	}

	log.Println("Successfully completed SASL negotiation!")

	return nil //errors.New("abort here")
}
