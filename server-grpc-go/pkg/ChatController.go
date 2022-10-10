package pkg

import (
    "context"
    "fmt"
    "github.com/golang/protobuf/ptypes/empty"
    interface_grpc "ningenme/application-interface-sample/server-grpc-go/interface-grpc"
    "os"
    "strconv"
    "time"
)

type ChatController struct {
    interface_grpc.UnimplementedChatServiceServer
}

var chatRepository = ChatRepository{}

var PollingMilliSecond = os.Getenv("POLLING_MILLISECOND")

func (c *ChatController) Post(ctx context.Context, req *interface_grpc.PostRequest) (*empty.Empty, error) {

    chatRepository.Insert(req.GetBody())

    return &empty.Empty{}, nil
}
func (c *ChatController) Get(empty *empty.Empty, stream interface_grpc.ChatService_GetServer) error {

    lastTime := time.Date(2022,1,1,0,0,0, 0, time.Local)
    ms, _ := strconv.Atoi(PollingMilliSecond)

    for {
        currentTime := time.Now()

        time.Sleep(time.Millisecond * time.Duration(ms))
        fmt.Println(lastTime, currentTime)

        list := chatRepository.Select(lastTime, currentTime)
        lastTime = currentTime

        fmt.Println(list)

        if len(list) == 0 {
            continue
        }
        if err := stream.Send(&interface_grpc.GetResponse{ChatList: list}); err != nil {
            return err
        }

    }
}

