#!/bin/bash
CURRENT_DIR=$1
for x in $(find ${CURRENT_DIR}/delever_protos/* -type d); do
  protoc -I=${x} -I=${CURRENT_DIR}/delever_protos -I /usr/local/include --go_out=plugins=grpc,strip_import_prefix=github.com/:${CURRENT_DIR} ${x}/*.proto
done
