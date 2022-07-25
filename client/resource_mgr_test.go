package client

import (
	"context"
	"github.com/kom0055/go-hadoop/proto/yarn/api"
	"github.com/kom0055/go-hadoop/proto/yarn/api/server"
	"testing"
)

func TestUpdateNodeResource(t *testing.T) {
	ctx := context.Background()
	rmClient, err := DialResourceManagerAdministrationProtocolService(ctx, "192.168.1.136:8141")
	if err != nil {
		t.Fatal(err)
	}
	memMB := int64(2048)
	core := int32(9)
	host := "office001134"
	port := int32(45454)
	resp, err := rmClient.UpdateNodeResource(ctx, &server.UpdateNodeResourceRequestProto{
		NodeResourceMap: []*api.NodeResourceMapProto{
			{
				NodeId: &api.NodeIdProto{
					Host: &host,
					Port: &port,
				},
				ResourceOption: &api.ResourceOptionProto{
					Resource: &api.ResourceProto{
						Memory:       &memMB,
						VirtualCores: &core,
					},
				},
			},
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

}
