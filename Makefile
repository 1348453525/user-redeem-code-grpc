GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)

.PHONY: tidy
# go mod tidy
tidy:
	go mod tidy

PROTO_DIR=./user-srv/proto
PROTO_FILES=$(shell find $(PROTO_DIR) -name *.proto)
PROTO_OUT_DIR=./user-srv/proto

.PHONY: pb
# 生成 proto 及 grpc 代码
pb:
	protoc \
	--proto_path=$(PROTO_DIR) \
	--go_out=$(PROTO_OUT_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(PROTO_OUT_DIR) \
	--go-grpc_opt=paths=source_relative \
	$(PROTO_FILES)

help:
	@echo "make help - 显示帮助信息"
	@echo "make tidy - go mod tidy"
	@echo "make pb - 生成 proto 及 grpc 代码"

.DEFAULT_GOAL := help
