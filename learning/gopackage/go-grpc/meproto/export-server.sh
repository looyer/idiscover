#!/bin/bash

# go install github.com/golang/protobuf/protoc-gen-go@latest
# https://github.com/protocolbuffers/protobuf/releases/tag/v22.0-rc1  protoc-22.0-rc-1-win64.zip

protoc --proto_path=./ --go_out=plugins=grpc:./ --go_opt=paths=source_relative common.proto 

