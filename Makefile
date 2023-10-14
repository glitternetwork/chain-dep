all: init proto-gen

proto-compiler:

init:

proto-gen:
	rm -rf glitter_proto/x/index/types/*.pb.go
	rm -rf glitter_proto/x/index/types/*.gw.go
	sh ./scripts/proto-gen.sh

.PHONY: all proto-gen init
