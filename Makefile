include base.proto.mk

HDP_DIR=${PROTO_ROOT_PATH}/github.com/kom0055/go-hadoop/proto/v1alpha1
HDP_COMMON_DIR=${HDP_DIR}/common

HDFS_DIR=${HDP_DIR}/hdfs
HDFS_CLIENT_DIR=${HDFS_DIR}/client
HDFS_RBF_DIR=${HDFS_DIR}/rbf
HDFS_DATANODE_DIR=${HDFS_DIR}/datanode
HDFS_NAMENODE_DIR=${HDFS_DIR}/namenode
HDFS_QJOURNAL_DIR=${HDFS_DIR}/qjournal

YARN_DIR=${HDP_DIR}/yarn
YARN_API_DIR=${YARN_DIR}/api
YARN_API_SERVER_DIR=${YARN_DIR}/api/server
YARN_COMMON_DIR=${YARN_DIR}/common
YARN_AM_DIR=${YARN_DIR}/am
YARN_SERVER_DIR=${YARN_DIR}/server
YARN_SERVER_COMMON_DIR=${YARN_DIR}/server/common
YARN_SERVER_APP_DIR=${YARN_DIR}/server/app
YARN_SERVER_NM_DIR=${YARN_DIR}/server/nm
YARN_SERVER_RM_DIR=${YARN_DIR}/server/rm



MR_DIR=${HDP_DIR}/mr
MR_CLIENT_DIR=${MR_DIR}/client
MR_CLIENT_COMMON_DIR=${MR_CLIENT_DIR}/common
MR_CLIENT_SHUFFLE_DIR=${MR_CLIENT_DIR}/shuffle

.PHONY: hdp-common
hdp-common:
	for file in $(HDP_COMMON_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH}; \
  	done



.PHONY: hdfs
hdfs: hdp-common
	for file in $(HDFS_CLIENT_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${HDP_COMMON_DIR} -I${HDFS_CLIENT_DIR}; \
  	done

	for file in $(HDFS_RBF_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${HDP_COMMON_DIR} -I${HDFS_CLIENT_DIR} -I${HDFS_RBF_DIR}; \
  	done

	for file in $(HDFS_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${HDFS_DIR} -I${HDP_COMMON_DIR} -I${HDFS_CLIENT_DIR} -I${HDFS_RBF_DIR}; \
  	done

	for file in $(HDFS_DATANODE_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${HDFS_DIR} -I${HDP_COMMON_DIR} -I${HDFS_CLIENT_DIR} -I${HDFS_RBF_DIR}; \
  	done

	for file in $(HDFS_NAMENODE_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${HDFS_DIR} -I${HDP_COMMON_DIR} -I${HDFS_CLIENT_DIR} -I${HDFS_RBF_DIR}; \
  	done

	for file in $(HDFS_QJOURNAL_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${HDFS_DIR} -I${HDP_COMMON_DIR} -I${HDFS_CLIENT_DIR} -I${HDFS_RBF_DIR}; \
  	done


.PHONY: yarn
yarn: hdp-common
	for file in $(YARN_API_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${YARN_API_DIR} -I${HDP_COMMON_DIR}; \
  	done

	for file in $(YARN_API_SERVER_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${YARN_API_SERVER_DIR} -I${YARN_API_DIR} -I${HDP_COMMON_DIR}; \
  	done

	for file in $(YARN_COMMON_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${YARN_COMMON_DIR} -I${YARN_API_DIR} -I${HDP_COMMON_DIR}; \
  	done

	for file in $(YARN_AM_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${YARN_COMMON_DIR} -I${YARN_API_DIR} -I${HDP_COMMON_DIR}; \
  	done

	for file in $(YARN_SERVER_COMMON_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${YARN_COMMON_DIR} -I${YARN_API_DIR} -I${HDP_COMMON_DIR}; \
  	done

	for file in $(YARN_SERVER_APP_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${YARN_SERVER_COMMON_DIR} -I${YARN_COMMON_DIR} -I${YARN_API_DIR} -I${HDP_COMMON_DIR}; \
  	done

	for file in $(YARN_SERVER_NM_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${YARN_SERVER_COMMON_DIR} -I${YARN_COMMON_DIR} -I${YARN_API_DIR} -I${HDP_COMMON_DIR}; \
  	done

	for file in $(YARN_SERVER_RM_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${YARN_SERVER_COMMON_DIR} -I${HDP_COMMON_DIR} -I${YARN_COMMON_DIR}  -I${YARN_API_DIR}; \
  	done


.PHONY: mr
mr: yarn
	for file in $(MR_CLIENT_COMMON_DIR)/*.proto; do \
  		protoc \
  		--proto_path=${PROTO_ROOT_PATH} \
		--go_out=${PROTO_ROOT_PATH} \
  		$${file} \
  		; \
  	done

	for file in $(MR_CLIENT_SHUFFLE_DIR)/*.proto; do \
  		protoc $${file}  --go_out=${PROTO_ROOT_PATH}  --go_opt=paths=source_relative  \
  		--proto_path=${PROTO_ROOT_PATH} \
  		-I${MR_CLIENT_SHUFFLE_DIR} -I${HDP_COMMON_DIR}  -I${YARN_API_DIR}; \
  	done

.PHONY: all
all: hdp-common hdfs yarn mr
