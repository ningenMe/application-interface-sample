package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
    interface_grpc "ningenme/application-interface-sample/server-grpc-go/interface-grpc"
    "ningenme/application-interface-sample/server-grpc-go/pkg"
)

func main() {
    port := "9020"
	lis, err := net.Listen("tcp", ":" + port)
	if err != nil {
		log.Fatalln(err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	{
        interface_grpc.RegisterChatServiceServer(s, &pkg.ChatController{})
	}

	if err := s.Serve(lis); err != nil {
		log.Fatalln(err)
	}
}
