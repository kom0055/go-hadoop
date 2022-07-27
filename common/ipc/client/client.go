package ipc

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/kom0055/go-hadoop/common/defined"
	"github.com/kom0055/go-hadoop/common/log"
	"github.com/kom0055/go-hadoop/proto/common"
	"github.com/nu7hatch/gouuid"
	"google.golang.org/protobuf/proto"
	"strings"
)

type Client struct {
	ClientId      *uuid.UUID
	Ugi           *common.UserInformationProto
	ServerAddress string
	TCPNoDelay    bool
}

func (c *Client) String() string {
	buf := bytes.NewBufferString("")
	_, _ = fmt.Fprint(buf, "<clientId:", c.ClientId)
	_, _ = fmt.Fprint(buf, ", server:", c.ServerAddress)
	_, _ = fmt.Fprint(buf, ">")
	return buf.String()
}

func (c *Client) Call(ctx context.Context, rpc *common.RequestHeaderProto, rpcRequest proto.Message, rpcResponse proto.Message) error {
	// Create connIdentity
	connectionId := connIdentity{user: *c.Ugi.RealUser, protocol: *rpc.DeclaringClassProtocolName, address: c.ServerAddress}

	// Get connection to server
	//log.Println("Connecting...", c)
	conn, err := c.getConnection(ctx, connectionId)
	if err != nil {
		return err
	}

	// Create call and send request
	rpcCall := call{callId: 0, procedure: rpc, request: rpcRequest, response: rpcResponse}
	err = conn.sendRequest(c.ClientId, &rpcCall)
	if err != nil {
		log.Warnf("sendRequest: %+v", err)
		connectionMgr.RmvConnection(connectionId)
		return err
	}

	// Read & return response
	err = c.readResponse(ctx, conn, &rpcCall)
	if err != nil {
		log.Warnf("readResponse: %+v", err)
		connectionMgr.RmvConnection(connectionId)
		return err
	}
	return nil
}

func (c *Client) readResponse(ctx context.Context, conn *connection, rpcCall *call) error {
	// Read first 4 bytes to get total-length
	var totalLength int32 = -1
	var totalLengthBytes [4]byte
	if _, err := conn.Read(totalLengthBytes[0:4]); err != nil {
		log.Warnf("readResponse#conn.con.Read(totalLengthBytes): %+v", err)
		return err
	}

	if err := defined.ConvertBytesToFixed(totalLengthBytes[0:4], &totalLength); err != nil {
		log.Warnf("readResponse#gohadoop.ConvertBytesToFixed(totalLengthBytes, &totalLength): %+v", err)
		return err
	}

	responseBytes := make([]byte, totalLength)
	if _, err := conn.Read(responseBytes); err != nil {
		log.Warnf("conn.con.Read(responseBytes): %+v", err)
		return err
	}

	// Parse RpcResponseHeaderProto
	rpcResponseHeaderProto := common.RpcResponseHeaderProto{}
	off, err := readDelimited(responseBytes[0:totalLength], &rpcResponseHeaderProto)
	if err != nil {
		log.Warnf("readDelimited(responseBytes, rpcResponseHeaderProto): %+v", err)
		return err
	}
	//log.Println("Received rpcResponseHeaderProto = ", rpcResponseHeaderProto)

	err = c.checkRpcHeader(ctx, &rpcResponseHeaderProto)
	if err != nil {
		log.Warnf("c.checkRpcHeader failed: %+v", err)
		return err
	}

	if *rpcResponseHeaderProto.Status == common.RpcResponseHeaderProto_SUCCESS {
		// Parse RpcResponseWrapper
		_, err = readDelimited(responseBytes[off:], rpcCall.response)
	} else {
		log.Warnf("RPC failed with status: %+s", rpcResponseHeaderProto.Status.String())
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

func (c *Client) checkRpcHeader(ctx context.Context, rpcResponseHeaderProto *common.RpcResponseHeaderProto) error {
	callClientId := *c.ClientId
	headerClientId := rpcResponseHeaderProto.ClientId
	if len(rpcResponseHeaderProto.ClientId) > 0 {
		if !bytes.Equal(callClientId[0:16], headerClientId[0:16]) {
			log.Infof("Incorrect clientId: %+v", headerClientId)
			return errors.New("incorrect clientId")
		}
	}
	return nil
}

func (c *Client) getConnection(ctx context.Context, connectionId connIdentity) (*connection, error) {
	// Try to re-use an existing connection

	conn, err := connectionMgr.GetConnection(ctx, c.ClientId, connectionId)

	// If necessary, create a new connection and save it in the connection-pool
	if err != nil {
		log.Warnf("Couldn't setup connection: %+v", err)
		return nil, err
	}

	return conn, nil
}
