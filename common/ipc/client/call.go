package ipc

import (
	"google.golang.org/protobuf/proto"
)

type call struct {
	callId     int32
	procedure  proto.Message
	request    proto.Message
	response   proto.Message
	retryCount int32
}
