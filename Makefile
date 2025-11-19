PROTOC_BIN := $(PWD)/tools/protoc/bin
PROTOC_NPM := $(PWD)/tools/protoc/node_modules/.bin
PROTOC_PY  := $(PWD)/tools/protoc/venv/bin

.PHONY: proto
proto:
	PATH="$(PROTOC_BIN):$(PROTOC_NPM):$(PROTOC_PY):$$PATH" && \
	cd proto && buf generate
