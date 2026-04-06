GOHOSTOS:=$(shell go env GOHOSTOS)
GOPATH:=$(shell go env GOPATH)

.PHONY: tidy
# go mod tidy
tidy:
	go mod tidy

USER_SRV_PROTO_DIR=./user-srv/proto/user
USER_SRV_PROTO_FILES=$(shell find $(USER_SRV_PROTO_DIR) -name *.proto)
USER_SRV_PROTO_OUT_DIR=./user-srv/proto/user

.PHONY: user-srv-proto
# 生成 user-srv 的 proto 及 grpc 代码
user-srv-proto:
	protoc \
	--proto_path=$(USER_SRV_PROTO_DIR) \
	--go_out=$(USER_SRV_PROTO_OUT_DIR) \
	--go_opt=paths=source_relative \
	--go-grpc_out=$(USER_SRV_PROTO_OUT_DIR) \
	--go-grpc_opt=paths=source_relative \
	$(USER_SRV_PROTO_FILES)

help:
	@echo "make help - 显示帮助信息"
	@echo "make tidy - go mod tidy"
	@echo "make user-srv-proto - 生成 user-srv 的 proto 及 grpc 代码"

.DEFAULT_GOAL := help
