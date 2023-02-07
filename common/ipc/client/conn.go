package ipc

import (
	"errors"
	"net"
	"strings"
	"time"

	"github.com/kom0055/go-hadoop/common/defined"
	"github.com/kom0055/go-hadoop/common/log"
	"github.com/kom0055/go-hadoop/common/security"
	"github.com/kom0055/go-hadoop/proto/v1alpha1/common"
	uuid "github.com/nu7hatch/gouuid"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

type connection struct {
	con net.Conn
}

func (conn *connection) Read(b []byte) (n int, err error) {
	return conn.con.Read(b)
}

func (conn *connection) Write(b []byte) (n int, err error) {
	return conn.con.Write(b)
}

func (conn *connection) Close() error {
	return conn.con.Close()
}

func (conn *connection) LocalAddr() net.Addr {
	return conn.con.LocalAddr()
}

func (conn *connection) RemoteAddr() net.Addr {
	return conn.con.RemoteAddr()
}

func (conn *connection) SetDeadline(t time.Time) error {
	return conn.con.SetDeadline(t)
}

func (conn *connection) SetReadDeadline(t time.Time) error {
	return conn.con.SetReadDeadline(t)
}

func (conn *connection) SetWriteDeadline(t time.Time) error {
	return conn.con.SetWriteDeadline(t)
}

func (conn *connection) writeConnectionHeader(authProtocol defined.AuthProtocol) error {
	// RPC_HEADER
	if _, err := conn.Write(defined.RpcHeader); err != nil {
		log.Warnf("conn.Write gohadoop.RPC_HEADER: %v", err)
		return err
	}

	// RPC_VERSION
	if _, err := conn.Write(defined.Version); err != nil {
		log.Warnf("conn.Write gohadoop.VERSION: %v", err)
		return err
	}

	// RPC_SERVICE_CLASS
	if serviceClass, err := defined.ConvertFixedToBytes(defined.RpcServiceClass); err != nil {
		log.Warnf("binary.Write: %v", err)
		return err
	} else if _, err := conn.Write(serviceClass); err != nil {
		log.Warnf("conn.Write RPC_SERVICE_CLASS: %v", err)
		return err
	}

	// AuthProtocol
	if authProtocolBytes, err := defined.ConvertFixedToBytes(authProtocol); err != nil {
		log.Warnf("WTF AUTH_PROTOCOL: %v", err)
		return err
	} else if _, err := conn.Write(authProtocolBytes); err != nil {
		log.Warnf("conn.Write gohadoop.AUTH_PROTOCOL: %v", err)
		return err
	}

	return nil
}

func (conn *connection) writeConnectionContext(clientUUId *uuid.UUID, connectionId connIdentity, _ defined.AuthProtocol) error {
	// Create hadoop_common.IpcConnectionContextProto
	ugi, _ := defined.CreateSimpleUGIProto()
	ipcCtxProto := common.IpcConnectionContextProto{UserInfo: ugi, Protocol: &connectionId.protocol}

	// Create RpcRequestHeaderProto
	callId := int32(-3)
	clientId := [16]byte(*clientUUId)

	/*if (authProtocol == gohadoop.AUTH_PROTOCOL_SASL) {
	  callId = SASL_RPC_CALL_ID
	}*/

	rpcReqHeaderProto := common.RpcRequestHeaderProto{RpcKind: &defined.RpcProtocolBufffer, RpcOp: &defined.RpcFinalPacket, CallId: &callId, ClientId: clientId[0:16], RetryCount: &defined.RpcDefaultRetryCount}

	rpcReqHeaderProtoBytes, err := proto.Marshal(&rpcReqHeaderProto)
	if err != nil {
		log.Warnf("proto.Marshal(&rpcReqHeaderProto): %v", err)
		return err
	}

	ipcCtxProtoBytes, _ := proto.Marshal(&ipcCtxProto)
	if err != nil {
		log.Warnf("proto.Marshal(&ipcCtxProto): %v", err)
		return err
	}

	totalLength := len(rpcReqHeaderProtoBytes) + sizeVarint(len(rpcReqHeaderProtoBytes)) + len(ipcCtxProtoBytes) + sizeVarint(len(ipcCtxProtoBytes))
	tLen := int32(totalLength)
	totalLengthBytes, err := defined.ConvertFixedToBytes(tLen)

	if err != nil {
		log.Warnf("ConvertFixedToBytes(totalLength): %v", err)
		return err
	} else if _, err := conn.Write(totalLengthBytes); err != nil {
		log.Warnf("conn.con.Write(totalLengthBytes): %v", err)
		return err
	}

	if err := conn.writeDelimitedBytes(rpcReqHeaderProtoBytes); err != nil {
		log.Warnf("writeDelimitedBytes(conn, rpcReqHeaderProtoBytes): %v", err)
		return err
	}
	if err := conn.writeDelimitedBytes(ipcCtxProtoBytes); err != nil {
		log.Warnf("writeDelimitedBytes(conn, ipcCtxProtoBytes): %v", err)
		return err
	}

	return nil
}

func (conn *connection) writeDelimitedBytes(data []byte) error {
	if _, err := conn.Write(protowire.AppendVarint([]byte{}, uint64(len(data)))); err != nil {
		log.Warnf("conn.con.Write(protowire.AppendVarint(uint64(len(data)))): %v", err)
		return err
	}

	//if _, err := conn.con.Write(proto.EncodeVarint(uint64(len(data)))); err != nil {
	//	log.Fatal("conn.con.Write(proto.EncodeVarint(uint64(len(data))))", err)
	//	return err
	//}
	if _, err := conn.Write(data); err != nil {
		log.Warnf("conn.con.Write(data): %v", err)
		return err
	}

	return nil
}

//nolint:golint,unused
func (conn *connection) writeDelimitedTo(msg proto.Message) error {
	msgBytes, err := proto.Marshal(msg)
	if err != nil {
		log.Warnf("proto.Marshal(msg): %v", err)
		return err
	}
	return conn.writeDelimitedBytes(msgBytes)
}

func (conn *connection) sendSaslMessage(message *common.RpcSaslProto) error {
	saslRpcHeaderProto := common.RpcRequestHeaderProto{RpcKind: &defined.RpcProtocolBufffer,
		RpcOp:      &defined.RpcFinalPacket,
		CallId:     &SaslRpcCallId,
		ClientId:   SaslRpcDummyClientId,
		RetryCount: &SaslRpcInvalidRetryCount}

	saslRpcHeaderProtoBytes, err := proto.Marshal(&saslRpcHeaderProto)

	if err != nil {
		log.Warnf("proto.Marshal(&saslRpcHeaderProto): %v", err)
		return err
	}

	saslRpcMessageProtoBytes, err := proto.Marshal(message)

	if err != nil {
		log.Warnf("proto.Marshal(saslMessage): %v", err)
		return err
	}

	totalLength := len(saslRpcHeaderProtoBytes) + sizeVarint(len(saslRpcHeaderProtoBytes)) + len(saslRpcMessageProtoBytes) + sizeVarint(len(saslRpcMessageProtoBytes))
	tLen := int32(totalLength)

	if totalLengthBytes, err := defined.ConvertFixedToBytes(tLen); err != nil {
		log.Warnf("ConvertFixedToBytes(totalLength): %v", err)
		return err
	} else {
		if _, err := conn.Write(totalLengthBytes); err != nil {
			log.Warnf("conn.con.Write(totalLengthBytes): %v", err)
			return err
		}
	}
	if err := conn.writeDelimitedBytes(saslRpcHeaderProtoBytes); err != nil {
		log.Warnf("writeDelimitedBytes(conn, saslRpcHeaderProtoBytes) :%v", err)
		return err
	}
	if err := conn.writeDelimitedBytes(saslRpcMessageProtoBytes); err != nil {
		log.Warnf("writeDelimitedBytes(conn, saslRpcMessageProtoBytes): %v", err)
		return err
	}

	return nil
}

func (conn *connection) receiveSaslMessage() (*common.RpcSaslProto, error) {
	// Read first 4 bytes to get total-length
	var totalLength int32 = -1
	var totalLengthBytes [4]byte

	if _, err := conn.Read(totalLengthBytes[0:4]); err != nil {
		log.Warnf("receiveSaslMessage#conn.con.Read(totalLengthBytes): %v", err)
		return nil, err
	}
	if err := defined.ConvertBytesToFixed(totalLengthBytes[0:4], &totalLength); err != nil {
		log.Warnf("receiveSaslMessage#gohadoop.ConvertBytesToFixed(totalLengthBytes, &totalLength): +v", err)
		return nil, err
	}

	responseBytes := make([]byte, totalLength)

	if _, err := conn.Read(responseBytes); err != nil {
		log.Warnf("receiveSaslMessage#conn.con.Read(responseBytes): %v", err)
		return nil, err
	}

	// Parse RpcResponseHeaderProto
	rpcResponseHeaderProto := common.RpcResponseHeaderProto{}
	off, err := readDelimited(responseBytes[0:totalLength], &rpcResponseHeaderProto)
	if err != nil {
		log.Warnf("readDelimited(responseBytes, rpcResponseHeaderProto): %v", err)
		return nil, err
	}

	err = checkSaslRpcHeader(&rpcResponseHeaderProto)
	if err != nil {
		log.Warnf("checkSaslRpcHeader failed: %v", err)
		return nil, err
	}

	var saslRpcMessage common.RpcSaslProto

	if *rpcResponseHeaderProto.Status == common.RpcResponseHeaderProto_SUCCESS {
		// Parse RpcResponseWrapper
		if _, err = readDelimited(responseBytes[off:], &saslRpcMessage); err != nil {
			log.Warnf("failed to read sasl response!")
			return nil, err
		} else {
			return &saslRpcMessage, nil
		}
	} else {
		log.Warnf("RPC failed with status: %s", rpcResponseHeaderProto.Status.String())
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

func (conn *connection) negotiateSimpleTokenAuth(addr string) error {
	var saslNegotiateState = common.RpcSaslProto_NEGOTIATE
	var saslNegotiateMessage = common.RpcSaslProto{State: &saslNegotiateState}
	var saslResponseMessage *common.RpcSaslProto
	var err error

	//send a SASL negotiation request
	if err = conn.sendSaslMessage(&saslNegotiateMessage); err != nil {
		log.Warnf("failed to send SASL NEGOTIATE message!")
		return err
	}

	//get a response with supported mehcanisms/challenge
	if saslResponseMessage, err = conn.receiveSaslMessage(); err != nil {
		log.Warnf("failed to receive SASL NEGOTIATE response!")
		return err
	}

	auths := saslResponseMessage.GetAuths()

	if numAuths := len(auths); numAuths <= 0 {
		log.Warnf("No supported auth mechanisms!")
		return errors.New("no supported auth mechanisms")
	}

	//for now, we only support auth when TOKEN/DIGEST-MD5 is the first/only
	//supported auth mechanism
	auth := auths[0]

	if !(auth.GetMethod() == "TOKEN" && auth.GetMechanism() == "DIGEST-MD5") {
		log.Warnf("gohadoop only supports TOKEN/DIGEST-MD5 auth!")
		return errors.New("gohadoop only supports TOKEN/DIGEST-MD5 auth")
	}

	method := auth.GetMethod()
	mechanism := auth.GetMechanism()
	protocol := auth.GetProtocol()
	serverId := auth.GetServerId()
	challenge := auth.GetChallenge()

	//TODO: token/service mapping + token selection based on type/service
	//we wouldn't have gotten this far if there wasn't at least one available token.
	userToken, _ := findUsableTokenForService(addr)
	response, err := security.GetDigestMD5ChallengeResponse(protocol, serverId, challenge, userToken)

	if err != nil {
		log.Warnf("failed to get challenge response! ", err)
		return err
	}

	saslInitiateState := common.RpcSaslProto_INITIATE
	authSend := common.RpcSaslProto_SaslAuth{Method: &method, Mechanism: &mechanism,
		Protocol: &protocol, ServerId: &serverId}
	authsSendArray := []*common.RpcSaslProto_SaslAuth{&authSend}
	saslInitiateMessage := common.RpcSaslProto{State: &saslInitiateState,
		Token: []byte(response), Auths: authsSendArray}

	//send a SASL inititate request
	if err = conn.sendSaslMessage(&saslInitiateMessage); err != nil {
		log.Warnf("failed to send SASL INITIATE message!")
		return err
	}

	//get a response with supported mehcanisms/challenge
	if saslResponseMessage, err = conn.receiveSaslMessage(); err != nil {
		log.Warnf("failed to read response to SASL INITIATE response!")
		return err
	}

	if saslResponseMessage.GetState() != common.RpcSaslProto_SUCCESS {
		log.Warnf("expected SASL SUCCESS response!")
		return errors.New("expected SASL SUCCESS response")
	}

	log.Infof("Successfully completed SASL negotiation!")

	return nil //errors.New("abort here")
}

func (conn *connection) sendRequest(clientId *uuid.UUID, rpcCall *call) error {
	//log.Println("About to call RPC: ", rpcCall.procedure)

	// 0. RpcRequestHeaderProto
	rpcReqHeaderProto := common.RpcRequestHeaderProto{RpcKind: &defined.RpcProtocolBufffer, RpcOp: &defined.RpcFinalPacket, CallId: &rpcCall.callId, ClientId: clientId[0:16], RetryCount: &rpcCall.retryCount}
	rpcReqHeaderProtoBytes, err := proto.Marshal(&rpcReqHeaderProto)
	if err != nil {
		log.Warnf("proto.Marshal(&rpcReqHeaderProto): %v", err)
		return err
	}

	// 1. RequestHeaderProto
	requestHeaderProto := rpcCall.procedure
	requestHeaderProtoBytes, err := proto.Marshal(requestHeaderProto)
	if err != nil {
		log.Warnf("proto.Marshal(&requestHeaderProto): %v", err)
		return err
	}

	// 2. Param
	paramProto := rpcCall.request
	paramProtoBytes, err := proto.Marshal(paramProto)
	if err != nil {
		log.Warnf("proto.Marshal(&paramProto): %v", err)
		return err
	}

	totalLength := len(rpcReqHeaderProtoBytes) + sizeVarint(len(rpcReqHeaderProtoBytes)) + len(requestHeaderProtoBytes) + sizeVarint(len(requestHeaderProtoBytes)) + len(paramProtoBytes) + sizeVarint(len(paramProtoBytes))
	tLen := int32(totalLength)
	if totalLengthBytes, err := defined.ConvertFixedToBytes(tLen); err != nil {
		log.Warnf("ConvertFixedToBytes(totalLength): %v", err)
		return err
	} else {
		if _, err := conn.Write(totalLengthBytes); err != nil {
			log.Warnf("conn.con.Write(totalLengthBytes): %v", err)
			return err
		}
	}

	if err := conn.writeDelimitedBytes(rpcReqHeaderProtoBytes); err != nil {
		log.Warnf("writeDelimitedBytes(conn, rpcReqHeaderProtoBytes): %v", err)
		return err
	}
	if err := conn.writeDelimitedBytes(requestHeaderProtoBytes); err != nil {
		log.Warnf("writeDelimitedBytes(conn, requestHeaderProtoBytes): %v", err)
		return err
	}
	if err := conn.writeDelimitedBytes(paramProtoBytes); err != nil {
		log.Warnf("writeDelimitedBytes(conn, paramProtoBytes): %v", err)
		return err
	}

	//log.Println("Succesfully sent request of length: ", totalLength)

	return nil
}
