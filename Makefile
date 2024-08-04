all: init proto-gen

proto-compiler:

init:

proto-gen:
	rm -rf glitter_proto/blockved/glitterchain/index/types/*.pb.go
	rm -rf glitter_proto/blockved/glitterchain/index/types/*.gw.go

	rm -rf glitter_proto/blockved/glitterchain/consumer/types/*.pb.go
	rm -rf glitter_proto/blockved/glitterchain/consumer/types/*.gw.go

	sh ./scripts/proto-gen.sh

.PHONY: all proto-gen init
