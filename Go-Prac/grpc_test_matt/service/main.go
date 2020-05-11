package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"
	pb "vincent/practice/Go-Prac/grpc_test_matt/pb"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type Hello struct {
}

func (h *Hello) Say(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {

	item := &pb.HelloReply{}

	item.Result = "Matt say: " + in.Message + in.Work + in.Eat
	fmt.Println(in)

	return item, nil
}

func (h *Hello) Saaaaay(stream pb.Hello_SaaaaayServer) error {
	stop := make(chan struct{}, 1)
	go func() {
		for {
			resp := &pb.HelloReply{
				Result: time.Now().String(),
			}
			stream.Send(resp)

			<-time.After(10 * time.Second)
		}
	}()

	go func() {
		for {
			in, err := stream.Recv()
			if err == io.EOF {
				log.Println("recv close")
				stop <- struct{}{}
			}
			if err != nil {
				log.Println("recv err:", err)
			}

			log.Println("in:", in)
		}
	}()

	select {
	case <-stop:
		return nil
	}
}

func main() {
	// 監聽指定埠口，這樣服務才能在該埠口執行。
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		fmt.Println("無法監聽該埠口：%v", err)
	}

	// 建立新 gRPC 伺服器並註冊服務。
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &Hello{})

	// 在 gRPC 伺服器上註冊反射服務。
	reflection.Register(s)

	// 開始在指定埠口中服務。
	if err := s.Serve(lis); err != nil {
		fmt.Println("無法提供服務：%v", err)
	}

}
