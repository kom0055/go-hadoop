package client

import (
	"context"
	"encoding/json"
	"math"

	"github.com/kom0055/go-hadoop/common/defined"
	ipc "github.com/kom0055/go-hadoop/common/ipc/client"
	"github.com/kom0055/go-hadoop/proto/v1alpha1/yarn/api"
	uuid "github.com/nu7hatch/gouuid"
	"google.golang.org/protobuf/proto"
)

var (
	_                         = proto.Marshal
	_                         = &json.SyntaxError{}
	_                         = math.Inf
	ApplicationClientProtocol = "org.apache.hadoop.yarn.api.ApplicationClientProtocolPB"
)

// DialApplicationClientProtocolService connects to an ApplicationClientProtocolService at the specified network address.
func DialApplicationClientProtocolService(ctx context.Context, serverAddress string) (*ApplicationClientProtocolServiceClient, error) {
	clientId, _ := uuid.NewV4()
	ugi, _ := defined.CreateSimpleUGIProto()
	c := &ipc.Client{ClientId: clientId, Ugi: ugi, ServerAddress: serverAddress}
	return &ApplicationClientProtocolServiceClient{c}, nil
}

type ApplicationClientProtocolServiceClient struct {
	*ipc.Client
}

func (c *ApplicationClientProtocolServiceClient) GetNewApplication(ctx context.Context,
	req *api.GetNewApplicationRequestProto) (*api.GetNewApplicationResponseProto, error) {
	resp := &api.GetNewApplicationResponseProto{}
	err := c.Call(ctx, defined.GetCalleeRPCRequestHeaderProto(&ApplicationClientProtocol), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ApplicationClientProtocolServiceClient) GetApplications(ctx context.Context,
	req *api.GetApplicationsRequestProto) (*api.GetApplicationsResponseProto, error) {
	resp := &api.GetApplicationsResponseProto{}
	err := c.Call(ctx, defined.GetCalleeRPCRequestHeaderProto(&ApplicationClientProtocol), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ApplicationClientProtocolServiceClient) SignalToContainer(ctx context.Context,
	req *api.SignalContainerRequestProto) (*api.SignalContainerResponseProto, error) {
	resp := &api.SignalContainerResponseProto{}
	err := c.Call(ctx, defined.GetCalleeRPCRequestHeaderProto(&ApplicationClientProtocol), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (c *ApplicationClientProtocolServiceClient) GetContainers(ctx context.Context,
	req *api.GetContainersRequestProto) (*api.GetContainersResponseProto, error) {
	resp := &api.GetContainersResponseProto{}
	err := c.Call(ctx, defined.GetCalleeRPCRequestHeaderProto(&ApplicationClientProtocol), req, resp)
	if err != nil {
		return nil, err
	}
	return resp, nil
}
