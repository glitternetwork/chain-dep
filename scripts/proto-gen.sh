#!/bin/bash
set -o errexit -o nounset -o pipefail
command -v shellcheck >/dev/null && shellcheck "$0"


protoc_gen_gocosmos() {
  if ! grep "github.com/gogo/protobuf => github.com/regen-network/protobuf" go.mod &>/dev/null ; then
    echo -e "\tPlease run this command from somewhere inside the cosmos-sdk folder."
    return 1
  fi

  go get github.com/regen-network/cosmos-proto/protoc-gen-gocosmos@latest 2>/dev/null
}

protoc_gen_gocosmos
OUT_DIR="./glitter_proto"
mkdir -p "$OUT_DIR"
echo "Processing glitter proto files ..."
SRC_DIR="./proto"

COMMON_DIR=${SRC_DIR}"/common"
protoc \
  --proto_path=${SRC_DIR} \
  -I="third_party/proto/" \
  --gocosmos_out=plugins=interfacetype+grpc:./glitter_proto \
  --grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true:./glitter_proto \
  $(find  $COMMON_DIR -path -prune -o -name '*.proto' -print0 | xargs -0)
mkdir -p   glitter_proto/common
cp -r glitter_proto/github.com/glitternetwork/chain-dep/glitter_proto/common/*.go glitter_proto/common


CHAIN_DIR=${SRC_DIR}"/x"
protoc \
  --proto_path=${SRC_DIR} \
  -I="third_party/proto/" \
  --gocosmos_out=plugins=interfacetype+grpc:./glitter_proto \
  --grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true:./glitter_proto \
  $(find  ${CHAIN_DIR} -path -prune -o -name '*.proto' -print0 | xargs -0)
mkdir -p   glitter_proto/x/index/types glitter_proto/engine
cp -r glitter_proto/github.com/glitternetwork/chain-dep/glitter_proto/x/index/types/*.go  glitter_proto/x/index/types/


ENGINE_DIR=${SRC_DIR}"/engine"
protoc \
  --proto_path=${SRC_DIR} \
  -I="third_party/proto/" \
  --gocosmos_out=plugins=interfacetype+grpc:./glitter_proto \
  --grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true:./glitter_proto \
  $(find  ${ENGINE_DIR} -path -prune -o -name '*.proto' -print0 | xargs -0)
mkdir -p   glitter_proto/engine
cp -r glitter_proto/github.com/glitternetwork/chain-dep/glitter_proto/engine/*.go glitter_proto/engine/


rm -rf glitter_proto/github.com/
