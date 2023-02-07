package trans

import (
	"fmt"

	"github.com/kom0055/go-hadoop/common/log"
	"github.com/kom0055/go-hadoop/proto/v1alpha1/yarn/api"
)

const (
	offset = 40
)

func GetContainerIdProto(containerId string) (*api.ContainerIdProto, error) {
	var clusterTs int64
	var appId, attemptId int32
	var epoch, cId int64
	if _, err := fmt.Sscanf(containerId, "container_e%d_%d_%d_%d_%d", &epoch, &clusterTs, &appId, &attemptId, &cId); err != nil {
		return nil, err
	}
	id := epoch<<offset + cId
	log.Infof("KillContainer# containerId: %v, clusterTs: %v, appId: %v, attemptId: %v, epoch: %v, cId: %v, id: %v",
		containerId, clusterTs, appId, attemptId, epoch, cId, id)

	applicationId := &api.ApplicationIdProto{
		Id:               &appId,
		ClusterTimestamp: &clusterTs,
	}
	appAttemptId := &api.ApplicationAttemptIdProto{
		ApplicationId: applicationId,
		AttemptId:     &attemptId,
	}
	containerIdProto := &api.ContainerIdProto{
		AppId:        applicationId,
		AppAttemptId: appAttemptId,
		Id:           &id,
	}

	return containerIdProto, nil
}
