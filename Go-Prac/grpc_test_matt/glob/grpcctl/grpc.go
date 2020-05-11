package grpcctl

import (
	"fmt"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var (
	Server *grpc.Server
)

func init() {
	// 監聽指定埠口，這樣服務才能在該埠口執行。
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("無法監聽該埠口：%v", err)
	}

	// 建立新 gRPC 伺服器並註冊服務。
	Server = grpc.NewServer()
	// hello.RegisterHelloServer(s, &Hello{})

	// 在 gRPC 伺服器上註冊反射服務。
	reflection.Register(Server)

	// 開始在指定埠口中服務。
	if err := Server.Serve(lis); err != nil {
		fmt.Println("無法提供服務：%v", err)
	}
}

// type RegisterServer func(server *grpc.Server, f interface{})
type RegisterServer interface {
}

func RegisterNewServer(rs RegisterServer) {

}
