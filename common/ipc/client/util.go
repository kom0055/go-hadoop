package ipc

import (
	"bytes"
	"errors"
	"github.com/kom0055/go-hadoop/common/log"
	"github.com/kom0055/go-hadoop/common/security"
	"github.com/kom0055/go-hadoop/proto/common"
	"google.golang.org/protobuf/encoding/protowire"
	"google.golang.org/protobuf/proto"
)

var (
	SaslRpcDummyClientId           = make([]byte, 0)
	SaslRpcCallId            int32 = -33
	SaslRpcInvalidRetryCount int32 = -1
)

func findUsableTokenForService(service string) (*common.TokenProto, bool) {
	userTokens := security.GetCurrentUser().GetUserTokens()

	log.Infof("looking for token for service: %s", service)

	if len(userTokens) == 0 {
		return nil, false
	}

	token := userTokens[service]
	if token != nil {
		return token, true
	}

	return nil, false
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

func readDelimited(rawData []byte, msg proto.Message) (int, error) {
	headerLength, off := protowire.ConsumeVarint(rawData)
	if off == 0 {
		log.Warnf("protowire.ConsumeVarint(rawData) returned zero")
		return -1, nil
	}

	//headerLength, off := proto.DecodeVarint(rawData)
	//if off == 0 {
	//	log.Fatal("proto.DecodeVarint(rawData) returned zero")
	//	return -1, nil
	//}

	err := proto.Unmarshal(rawData[off:off+int(headerLength)], msg)
	if err != nil {
		log.Warnf("proto.Unmarshal(rawData[off:off+headerLength]): %+v ", err)
		return -1, err
	}

	return off + int(headerLength), nil
}

func checkSaslRpcHeader(rpcResponseHeaderProto *common.RpcResponseHeaderProto) error {
	headerClientId := rpcResponseHeaderProto.ClientId
	if rpcResponseHeaderProto.ClientId != nil {
		if !bytes.Equal(SaslRpcDummyClientId, headerClientId) {
			log.Warnf("Incorrect clientId: %+v", headerClientId)
			return errors.New("incorrect clientId")
		}
	}
	return nil
}
