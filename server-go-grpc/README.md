# server-go-grpc

```shell
grpcurl -plaintext localhost:${AIS_SERVER_GRPC_GO_PORT} list
grpcurl -plaintext localhost:${AIS_SERVER_GRPC_GO_PORT} list sample.ChatService
grpcurl -plaintext localhost:${AIS_SERVER_GRPC_GO_PORT} sample.ChatService.Get
grpcurl -plaintext -d '
    {
      "body" : "てすとめっせーじ"
    }
' localhost:${AIS_SERVER_GRPC_GO_PORT} sample.ChatService.Post
```
