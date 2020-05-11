package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"
	pb "vincent/practice/Go-Prac/grpc_test_matt/pb"

	"google.golang.org/grpc"
)

func main() {
	// 連線到遠端 gRPC 伺服器。
	conn, err := grpc.Dial("127.0.0.1:50051", grpc.WithInsecure())
	if err != nil {
		fmt.Println("連線失敗：%v", err)
	}
	defer conn.Close()
	// 建立新的 Calculator 客戶端，所以等一下就能夠使用 Calculator 的所有方法。
	c := pb.NewHelloClient(conn)

	// 傳送新請求到遠端 gRPC 伺服器 Calculator 中，並呼叫 Plus 函式，讓兩個數字相加。
	//Say(c)
	Saaaaay(c)
}

func Say(c pb.HelloClient) {
	cond := &pb.HelloRequest{}
	cond.Message = "that's life"
	cond.Eat = "pie"
	cond.Work = "lumberjack"

	foods := make([]*pb.Food, 0)
	foods = append(foods, &pb.Food{Key: "aaa", Val: "AAA"})
	foods = append(foods, &pb.Food{Key: "bbb", Val: "BBB"})
	cond.Food = foods
	r, err := c.Say(context.Background(), cond)
	if err != nil {
		fmt.Println("無法執行 Plus 函式：%v", err)
	}
	fmt.Println("回傳結果：", r.Result)
}

func Saaaaay(c pb.HelloClient) {
	stop := make(chan struct{}, 1)
	stream, err := c.Saaaaay(context.TODO())
	if err == io.EOF {
		stop <- struct{}{}
		//return
	} else if err != nil {
		log.Println(err)
		stop <- struct{}{}
	}

	go func() {
		for {
			resp := &pb.HelloRequest{
				Message: "client send: " + time.Now().String(),
			}
			stream.Send(resp)

			<-time.After(7 * time.Second)
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
		log.Println("close")
	}
}
