all: init proto-gen

proto-compiler:

init:

proto-gen:
	git submodule update
	rm -rf glitter_proto/glitterchain/index/types/*.pb.go
	rm -rf glitter_proto/glitterchain/index/types/*.gw.go

	rm -rf glitter_proto/glitterchain/consumer/types/*.pb.go
	rm -rf glitter_proto/glitterchain/consumer/types/*.gw.go

	sh ./scripts/proto-gen.sh

.PHONY: all proto-gen init
