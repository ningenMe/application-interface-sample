package pkg

import (
    "context"
    "fmt"
    "github.com/golang/protobuf/ptypes/empty"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    interface_grpc "ningenme/application-interface-sample/server-grpc-go/interface-grpc"
)

type ChatController struct {
    interface_grpc.UnimplementedChatServiceServer
}


func (c *ChatController) Post(context.Context, *interface_grpc.PostRequest) (*empty.Empty, error) {
    fmt.Println("Post")
    return nil, status.Errorf(codes.Unimplemented, "method Post not implemented")
}
func (c *ChatController) Get(empty *empty.Empty, stream interface_grpc.ChatService_GetServer) error {
    fmt.Println("Get")
    return nil
}

