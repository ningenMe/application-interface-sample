package pkg

import (
    "context"
    "fmt"
    "github.com/golang/protobuf/ptypes/empty"
    interface_grpc "ningenme/application-interface-sample/server-grpc-go/interface-grpc"
)

type ChatController struct {
    interface_grpc.UnimplementedChatServiceServer
}

var chatRepository = ChatRepository{}

func (c *ChatController) Post(ctx context.Context, req *interface_grpc.PostRequest) (*empty.Empty, error) {

    chatRepository.Insert(req.GetBody())

    return &empty.Empty{}, nil
}
func (c *ChatController) Get(empty *empty.Empty, stream interface_grpc.ChatService_GetServer) error {
    fmt.Println("Get")
    return nil
}

