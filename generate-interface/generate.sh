#!/bin/sh
protoc --version

GO_OUT_PATH="./server-grpc-go"
protoc \
--go_out="${GO_OUT_PATH}" --go_opt=paths=source_relative \
--go-grpc_out="${GO_OUT_PATH}" --go-grpc_opt=paths=source_relative \
./interface-grpc/*.proto

TS_OUT_PATH="./client-react/src"
protoc \
--js_out="import_style=commonjs,binary:${TS_OUT_PATH}" \
--grpc-web_out="import_style=typescript,mode=grpcwebtext:${TS_OUT_PATH}" \
./interface-grpc/*.proto
