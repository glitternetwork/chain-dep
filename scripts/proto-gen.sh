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
SRC_DIR="./submodules/glitter.proto/proto"


CHAIN_DIR=${SRC_DIR}"/glitterchain/index"
rm -rf  glitter_proto/glitterchain/index/types/*.pb.go
rm -rf  glitter_proto/glitterchain/index/types/*.pb.gw.go
protoc \
  --proto_path=${SRC_DIR} \
  --gocosmos_out=plugins=interfacetype+grpc:./glitter_proto \
  --grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true:./glitter_proto \
  $(find  ${CHAIN_DIR} -path -prune -o -name '*.proto' -print0 | xargs -0)
cp -r glitter_proto/github.com/glitternetwork/chain-dep/glitter_proto/glitterchain/index/types/*.go  glitter_proto/glitterchain/index/types/


CHAIN_DIR=${SRC_DIR}"/glitterchain/consumer"
protoc \
  --proto_path=${SRC_DIR} \
  --gocosmos_out=plugins=interfacetype+grpc:./glitter_proto \
  --grpc-gateway_out=logtostderr=true,allow_colon_final_segments=true:./glitter_proto \
  $(find  ${CHAIN_DIR} -path -prune -o -name '*.proto' -print0 | xargs -0)
cp -r glitter_proto/github.com/glitternetwork/chain-dep/glitter_proto/glitterchain/consumer/types/*.go  glitter_proto/glitterchain/consumer/types/




rm -rf glitter_proto/github.com/
