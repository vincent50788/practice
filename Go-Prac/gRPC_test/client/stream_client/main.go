package main

import (
	"context"
	"io"
	"log"
	"sync"

	"google.golang.org/grpc"

	pb "vincent/practice/Go-Prac/gRPC_test/pb"
)

const (
	PORT = "9002"
)

func main() {
	conn, err := grpc.Dial(":"+PORT, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	client := pb.NewStreamServiceClient(conn)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		err = printLists(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "first request", Value: 1010}})
		if err != nil {
			log.Fatalf("printLists.err: %v", err)
		}
		wg.Done()
	}()

	go func() {
		err = printLists(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "second request", Value: 2020}})
		if err != nil {
			log.Fatalf("printLists.err: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()

	//err = printRecord(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Record", Value: 2018}})
	//if err != nil {
	//	log.Fatalf("printRecord.err: %v", err)
	//}
	//
	//err = printRoute(client, &pb.StreamRequest{Pt: &pb.StreamPoint{Name: "gRPC Stream Client: Route", Value: 2018}})
	//if err != nil {
	//	log.Fatalf("printRoute.err: %v", err)
	//}
}

func printLists(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.List(context.Background(), r)
	if err != nil {
		return err
	}

	for {
		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}

	return nil
}

func printRecord(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Record(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n < 6; n++ {
		err := stream.Send(r)
		if err != nil {
			return err
		}
	}

	resp, err := stream.CloseAndRecv()
	if err != nil {
		return err
	}

	log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)

	return nil
}

func printRoute(client pb.StreamServiceClient, r *pb.StreamRequest) error {
	stream, err := client.Route(context.Background())
	if err != nil {
		return err
	}

	for n := 0; n <10; n++ {
		err = stream.Send(r)
		if err != nil {
			return err
		}

		resp, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}

		log.Printf("resp: pj.name: %s, pt.value: %d", resp.Pt.Name, resp.Pt.Value)
	}

	stream.CloseSend()

	return nil
}