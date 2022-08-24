package trans_test

import (
	"fmt"
	"testing"

	"github.com/kom0055/go-hadoop/common/trans"
)

func TestSscanf(t *testing.T) {
	var clusterTs int64
	var appId, attemptId int32
	var epoch, cId int64
	containerId := "container_e02_1655197268537_0004_01_000001"
	_, err := fmt.Sscanf(containerId, "container_e%d_%d_%d_%d_%d", &epoch, &clusterTs, &appId, &attemptId, &cId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(clusterTs, appId, attemptId, epoch, cId)
}

func TestGetContainerIdProto(t *testing.T) {

	containerId := "container_e02_1655197268537_0004_01_000001"

	cid, err := trans.GetContainerIdProto(containerId)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(cid)
}
