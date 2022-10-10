# application-interface-sample

```shell
GO_OUT_PATH="./server-grpc-go"
protoc \
--go_out="${GO_OUT_PATH}" --go_opt=paths=source_relative \
--go-grpc_out="${GO_OUT_PATH}" --go-grpc_opt=paths=source_relative \
./interface-grpc/*.proto
```

```shell
docker-compose -f docker-compose-generate-interface.yml build --no-cache &&  docker-compose -f docker-compose-generate-interface.yml up 
```