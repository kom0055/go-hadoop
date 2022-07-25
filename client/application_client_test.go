package client

import (
	"context"
	"github.com/kom0055/go-hadoop/proto/yarn/api"
	"testing"
)

const (
	ContainerIdBitmask = 0xffffffffff
)

func TestSignalToContainer(t *testing.T) {
	ctx := context.Background()
	appClient, err := DialApplicationClientProtocolService(ctx, "192.168.1.136:8050")
	if err != nil {
		t.Fatal(err)
	}
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
	ctx := context.Background()
	appClient, err := DialApplicationClientProtocolService(ctx, "192.168.1.136:8050")
	if err != nil {
		t.Fatal(err)
	}
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
	ctx := context.Background()
	appClient, err := DialApplicationClientProtocolService(ctx, "192.168.1.136:8050")
	if err != nil {
		t.Fatal(err)
	}
	limit := int64(5)
	resp, err := appClient.GetApplications(ctx, &api.GetApplicationsRequestProto{
		ApplicationTypes:  nil,
		ApplicationStates: nil,
		Users:             nil,
		Queues:            nil,
		Limit:             &limit,
		StartBegin:        nil,
		StartEnd:          nil,
		FinishBegin:       nil,
		FinishEnd:         nil,
		ApplicationTags:   nil,
		Scope:             nil,
	})
	if err != nil {
		t.Fatal(err)
	}
	t.Log(resp)

}
