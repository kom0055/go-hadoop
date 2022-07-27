package application_client_test_test

import (
	"context"
	"github.com/kom0055/go-hadoop/client"
	"github.com/kom0055/go-hadoop/common/log"
	"github.com/kom0055/go-hadoop/proto/yarn/api"
	"testing"
)

const (
	ContainerIdBitmask = 0xffffffffff
)

var (
	ctx       = context.Background()
	appClient *client.ApplicationClientProtocolServiceClient
)

func TestMain(m *testing.M) {
	var err error
	appClient, err = client.DialApplicationClientProtocolService(ctx, "192.168.1.136:8050")
	if err != nil {
		log.Fatalf("DialApplicationClientProtocolService FAILED: %+V", err)
	}
	m.Run()
}

func TestGetNewApplication(t *testing.T) {
	resp, err := appClient.GetNewApplication(ctx, &api.GetNewApplicationRequestProto{})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)
}

func TestSignalToContainer(t *testing.T) {
	appId := int32(1431)
	attemptId := int32(000001)
	clusterTs := int64(1658419814772)
	containerId := int64(5<<40 + 000006)
	applicationId := &api.ApplicationIdProto{
		Id:               &appId,
		ClusterTimestamp: &clusterTs,
	}
	appAttemptId := &api.ApplicationAttemptIdProto{
		ApplicationId: applicationId,
		AttemptId:     &attemptId,
	}
	cmd := api.SignalContainerCommandProto_GRACEFUL_SHUTDOWN
	resp, err := appClient.SignalToContainer(ctx, &api.SignalContainerRequestProto{
		ContainerId: &api.ContainerIdProto{
			AppId:        applicationId,
			AppAttemptId: appAttemptId,
			Id:           &containerId,
		},
		Command: &cmd,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

}

func TestGetContainers(t *testing.T) {

	appId := int32(1431)
	attemptId := int32(000001)
	clusterTs := int64(1658419814772)
	applicationId := &api.ApplicationIdProto{
		Id:               &appId,
		ClusterTimestamp: &clusterTs,
	}
	appAttemptId := &api.ApplicationAttemptIdProto{
		ApplicationId: applicationId,
		AttemptId:     &attemptId,
	}
	resp, err := appClient.GetContainers(ctx, &api.GetContainersRequestProto{
		ApplicationAttemptId: appAttemptId,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

}

func TestGetApplications(t *testing.T) {

	limit := int64(5)
	resp, err := appClient.GetApplications(ctx, &api.GetApplicationsRequestProto{
		Limit: &limit,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

}
