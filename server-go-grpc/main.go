package main

import (
    "github.com/jmoiron/sqlx"
    "google.golang.org/grpc"
    "google.golang.org/grpc/reflection"
    "log"
    "net"
    interface_grpc "ningenme/application-interface-sample/server-go-grpc/interface-grpc"
    "ningenme/application-interface-sample/server-go-grpc/pkg"
    "os"
)

var (
    ServerGrpcGoPort = os.Getenv("AIS_SERVER_GO_GRPC_PORT")
)

func main() {
	l, err := net.Listen("tcp", ":" +ServerGrpcGoPort)
	if err != nil {
		log.Fatalln(err)
	}

    pkg.DatabaseMysql, err = sqlx.Open("mysql", pkg.GetMysqlConfig())
    if err != nil {
        log.Fatalln(err)
    }
    defer pkg.DatabaseMysql.Close()

    s := grpc.NewServer()
	reflection.Register(s)
    interface_grpc.RegisterChatServiceServer(s, &pkg.ChatController{})

	if err := s.Serve(l); err != nil {
		log.Fatalln(err)
	}
}
