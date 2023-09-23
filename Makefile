all: init proto-gen

proto-compiler:

init:

proto-gen:
	sh ./scripts/proto-gen.sh

.PHONY: all proto-gen init
