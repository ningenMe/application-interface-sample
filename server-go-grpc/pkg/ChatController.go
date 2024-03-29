package pkg

import (
    "context"
    "fmt"
    "github.com/golang/protobuf/ptypes/empty"
    interface_grpc "ningenme/application-interface-sample/server-go-grpc/interface-grpc"
    "os"
    "strconv"
    "time"
)

type ChatController struct {
    interface_grpc.UnimplementedChatServiceServer
}

var chatRepository = ChatRepository{}

var PollingMilliSecond = os.Getenv("AIS_POLLING_MILLISECOND")

func (c *ChatController) Post(ctx context.Context, req *interface_grpc.PostRequest) (*empty.Empty, error) {

    chatRepository.Insert(req.GetBody())

    return &empty.Empty{}, nil
}
func (c *ChatController) Get(empty *empty.Empty, stream interface_grpc.ChatService_GetServer) error {

    lastTime := time.Date(2022,1,1,0,0,0, 0, time.Local)
    ms, _ := strconv.Atoi(PollingMilliSecond)

    for {
        if err := stream.Context().Err(); err != nil {
            fmt.Println(err)
            return nil
        }

        currentTime := time.Now()

        time.Sleep(time.Millisecond * time.Duration(ms))
        fmt.Println(lastTime, currentTime)

        list := chatRepository.Select(lastTime, currentTime)
        lastTime = currentTime

        if len(list) == 0 {
            continue
        }

        if err := stream.Send(&interface_grpc.GetResponse{ChatList: list}); err != nil {
            return err
        }

    }
}

