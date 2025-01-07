#!/bin/bash

CURRENT_DIR=$1
mkdir -p ${CURRENT_DIR}/doc/swagger

for x in $(find ${CURRENT_DIR}/protos -name '*.proto'); do
  if [[ $x == *"/google/"* ]] || [[ $x == *"/protoc-gen-openapiv2/"* ]]; then
    continue
  fi
  protoc \
    -I=${CURRENT_DIR}/protos \
    --go_out=${CURRENT_DIR} \
    --go_opt=module=milliy \
    --go-grpc_out=${CURRENT_DIR} \
    --go-grpc_opt=module=milliy \
    --grpc-gateway_out=${CURRENT_DIR} \
    --grpc-gateway_opt=module=milliy \
    --openapiv2_out=${CURRENT_DIR}/doc/swagger \
    --openapiv2_opt=allow_merge=true \
    --openapiv2_opt=merge_file_name=swagger_docs \
    $x
done
