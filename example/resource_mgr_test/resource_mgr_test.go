package resource_mgr_test_test

import (
	"context"
	"github.com/kom0055/go-hadoop/client"
	"github.com/kom0055/go-hadoop/proto/yarn/api"
	"github.com/kom0055/go-hadoop/proto/yarn/api/server"
	"log"
	"testing"
)

var (
	ctx      = context.Background()
	rmClient *client.ResourceManagerAdministrationProtocolServiceClient
)

func TestMain(m *testing.M) {
	var err error
	rmClient, err = client.DialResourceManagerAdministrationProtocolService(ctx, "192.168.1.136:8141")
	if err != nil {
		log.Fatalln(err)
	}
	m.Run()
}

func TestUpdateNodeResource(t *testing.T) {

	memMB := int64(10240)
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
