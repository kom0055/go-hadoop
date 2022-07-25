package client

import (
	"context"
	"encoding/json"
	"github.com/kom0055/go-hadoop/common/defined"
	ipc "github.com/kom0055/go-hadoop/common/ipc/client"
	"github.com/kom0055/go-hadoop/proto/yarn/api/server"
	uuid "github.com/nu7hatch/gouuid"
	"google.golang.org/protobuf/proto"
	"math"
)

var (
	_                                     = proto.Marshal
	_                                     = &json.SyntaxError{}
	_                                     = math.Inf
	ResourceManagerAdministrationProtocol = "org.apache.hadoop.yarn.server.api.ResourceManagerAdministrationProtocolPB"
)

// DialResourceManagerAdministrationProtocolService connects to an ApplicationClientProtocolService at the specified network address.
func DialResourceManagerAdministrationProtocolService(ctx context.Context, serverAddress string) (*ResourceManagerAdministrationProtocolServiceClient, error) {
	clientId, _ := uuid.NewV4()
	ugi, _ := defined.CreateSimpleUGIProto()
	c := &ipc.Client{ClientId: clientId, Ugi: ugi, ServerAddress: serverAddress}
	return &ResourceManagerAdministrationProtocolServiceClient{c}, nil
}

type ResourceManagerAdministrationProtocolServiceClient struct {
	*ipc.Client
}

func (c *ResourceManagerAdministrationProtocolServiceClient) UpdateNodeResource(ctx context.Context,
	req *server.UpdateNodeResourceRequestProto) (*server.UpdateNodeResourceResponseProto, error) {
	resp := &server.UpdateNodeResourceResponseProto{}
	err := c.Call(ctx, defined.GetCalleeRPCRequestHeaderProto(&ResourceManagerAdministrationProtocol), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
