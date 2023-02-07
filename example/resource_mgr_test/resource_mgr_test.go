package resource_mgr_test_test

import (
	"context"
	"log"
	"testing"
	"time"

	"github.com/kom0055/go-hadoop/client"
	"github.com/kom0055/go-hadoop/proto/v1alpha1/yarn/api"
	"github.com/kom0055/go-hadoop/proto/v1alpha1/yarn/api/server"
)

var (
	ctx      = context.Background()
	rmClient *client.ResourceManagerAdministrationProtocolServiceClient
)

func TestMain(m *testing.M) {
	var err error
	rmClient, err = client.DialResourceManagerAdministrationProtocolService(ctx, "192.168.1.111:8141")
	if err != nil {
		log.Fatalf("DialResourceManagerAdministrationProtocolService FAILED: %v", err)
	}
	m.Run()
}

func TestUpdateNodeResource(t *testing.T) {
	for {
		memMB := int64(16 * 1024)
		core := int32(80)
		host := "office001125"
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
			t.Log(err)
		}
		t.Log(resp)
		time.Sleep(time.Second * 2)

	}

}
