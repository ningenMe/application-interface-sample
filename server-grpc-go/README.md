# server-grpc-go

```shell
grpcurl -plaintext localhost:9140 list
grpcurl -plaintext localhost:9140 list sample.ChatService
grpcurl -plaintext localhost:9140 sample.ChatService.Get
grpcurl -plaintext -d '
    {
      "body" : "てすとめっせーじ"
    }
' localhost:9140 sample.ChatService.Post
```
