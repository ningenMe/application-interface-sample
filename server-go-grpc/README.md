# server-go-grpc

```shell
grpcurl -plaintext localhost:${AIS_SERVER_GO_GRPC_PORT} list
grpcurl -plaintext localhost:${AIS_SERVER_GO_GRPC_PORT} list sample.ChatService
grpcurl -plaintext localhost:${AIS_SERVER_GO_GRPC_PORT} sample.ChatService.Get
grpcurl -plaintext -d '
    {
      "body" : "てすとめっせーじ"
    }
' localhost:${AIS_SERVER_GO_GRPC_PORT} sample.ChatService.Post
```
