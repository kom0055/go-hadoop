PROTO_PATH=${GOPATH}/src
PROJECT_PATH=${PROTO_PATH}/github.com/kom0055/go-hadoop

COMMON_DIR=${PROJECT_PATH}/proto/common


HDFS_DIR=${PROJECT_PATH}/proto/hdfs
HDFS_CLIENT_DIR=${HDFS_DIR}/client
HDFS_RBF_DIR=${HDFS_DIR}/rbf
HDFS_DATANODE_DIR=${HDFS_DIR}/datanode
HDFS_NAMENODE_DIR=${HDFS_DIR}/namenode
HDFS_QJOURNAL_DIR=${HDFS_DIR}/qjournal


YARN_DIR=${PROJECT_PATH}/proto/yarn
YARN_API_DIR=${YARN_DIR}/api
YARN_API_SERVER_DIR=${YARN_DIR}/api/server
YARN_COMMON_DIR=${YARN_DIR}/common
YARN_AM_DIR=${YARN_DIR}/am
YARN_SERVER_DIR=${YARN_DIR}/server
YARN_SERVER_COMMON_DIR=${YARN_DIR}/server/common
YARN_SERVER_APP_DIR=${YARN_DIR}/server/app
YARN_SERVER_NM_DIR=${YARN_DIR}/server/nm
YARN_SERVER_RM_DIR=${YARN_DIR}/server/rm



MR_DIR=${PROJECT_PATH}/proto/mr
MR_CLIENT_DIR=${MR_DIR}/client
MR_CLIENT_COMMON_DIR=${MR_CLIENT_DIR}/common
MR_CLIENT_SHUFFLE_DIR=${MR_CLIENT_DIR}/shuffle

.PHONY: common
common:
	for file in $(COMMON_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${COMMON_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${COMMON_DIR}:${PROTO_PATH}; \
  	done


.PHONY: hdfs
hdfs: common
	for file in $(HDFS_CLIENT_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${HDFS_CLIENT_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${HDFS_CLIENT_DIR}:${PROTO_PATH} \
  		-I${COMMON_DIR} -I${HDFS_CLIENT_DIR}; \
  	done

	for file in $(HDFS_RBF_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${HDFS_RBF_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${HDFS_RBF_DIR}:${PROTO_PATH} \
  		-I${COMMON_DIR} -I${HDFS_CLIENT_DIR} -I${HDFS_RBF_DIR}; \
  	done

	for file in $(HDFS_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${HDFS_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${HDFS_DIR}:${PROTO_PATH} \
  		-I${HDFS_DIR} -I${COMMON_DIR} -I${HDFS_CLIENT_DIR} -I${HDFS_RBF_DIR}; \
  	done

	for file in $(HDFS_DATANODE_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${HDFS_DATANODE_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${HDFS_DATANODE_DIR}:${PROTO_PATH} \
  		-I${HDFS_DIR} -I${COMMON_DIR} -I${HDFS_CLIENT_DIR} -I${HDFS_RBF_DIR}; \
  	done

	for file in $(HDFS_NAMENODE_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${HDFS_NAMENODE_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${HDFS_NAMENODE_DIR}:${PROTO_PATH} \
  		-I${HDFS_DIR} -I${COMMON_DIR} -I${HDFS_CLIENT_DIR} -I${HDFS_RBF_DIR}; \
  	done

	for file in $(HDFS_QJOURNAL_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${HDFS_QJOURNAL_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${HDFS_QJOURNAL_DIR}:${PROTO_PATH} \
  		-I${HDFS_DIR} -I${COMMON_DIR} -I${HDFS_CLIENT_DIR} -I${HDFS_RBF_DIR}; \
  	done


.PHONY: yarn
yarn: common
	for file in $(YARN_API_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${YARN_API_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${YARN_API_DIR}:${PROTO_PATH} \
  		-I${YARN_API_DIR} -I${COMMON_DIR}; \
  	done

	for file in $(YARN_API_SERVER_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${YARN_API_SERVER_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${YARN_API_SERVER_DIR}:${PROTO_PATH} \
  		-I${YARN_API_SERVER_DIR} -I${YARN_API_DIR} -I${COMMON_DIR}; \
  	done

	for file in $(YARN_COMMON_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${YARN_COMMON_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${YARN_COMMON_DIR}:${PROTO_PATH} \
  		-I${YARN_COMMON_DIR} -I${YARN_API_DIR} -I${COMMON_DIR}; \
  	done

	for file in $(YARN_AM_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${YARN_AM_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${YARN_AM_DIR}:${PROTO_PATH} \
  		-I${YARN_COMMON_DIR} -I${YARN_API_DIR} -I${COMMON_DIR}; \
  	done

	for file in $(YARN_SERVER_COMMON_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${YARN_SERVER_COMMON_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${YARN_SERVER_COMMON_DIR}:${PROTO_PATH} \
  		-I${YARN_COMMON_DIR} -I${YARN_API_DIR} -I${COMMON_DIR}; \
  	done

	for file in $(YARN_SERVER_APP_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${YARN_SERVER_APP_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${YARN_SERVER_APP_DIR}:${PROTO_PATH} \
  		-I${YARN_SERVER_COMMON_DIR} -I${YARN_COMMON_DIR} -I${YARN_API_DIR} -I${COMMON_DIR}; \
  	done

	for file in $(YARN_SERVER_NM_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${YARN_SERVER_NM_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${YARN_SERVER_NM_DIR}:${PROTO_PATH} \
  		-I${YARN_SERVER_COMMON_DIR} -I${YARN_COMMON_DIR} -I${YARN_API_DIR} -I${COMMON_DIR}; \
  	done

	for file in $(YARN_SERVER_RM_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${YARN_SERVER_RM_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${YARN_SERVER_RM_DIR}:${PROTO_PATH} \
  		-I${YARN_SERVER_COMMON_DIR} -I${COMMON_DIR} -I${YARN_COMMON_DIR}  -I${YARN_API_DIR}; \
  	done





.PHONY: mr
mr: yarn
	for file in $(MR_CLIENT_COMMON_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${MR_CLIENT_COMMON_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${MR_CLIENT_COMMON_DIR}:${PROTO_PATH} \
  		-I${MR_CLIENT_COMMON_DIR} -I${COMMON_DIR}  -I${YARN_API_DIR}; \
  	done

	for file in $(MR_CLIENT_SHUFFLE_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${MR_CLIENT_SHUFFLE_DIR}  --go_opt=paths=source_relative  \
  		--proto_path=${MR_CLIENT_SHUFFLE_DIR}:${PROTO_PATH} \
  		-I${MR_CLIENT_SHUFFLE_DIR} -I${COMMON_DIR}  -I${YARN_API_DIR}; \
  	done




.PHONY: all
all: common hdfs yarn mr