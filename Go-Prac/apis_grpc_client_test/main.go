package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"sync"
	"time"
	"vincent/practice/Go-Prac/apis_grpc_client_test/zqbrecord"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		conn, err := grpc.Dial(":"+"17802", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("grpc.Dial err: %v", err)
		}

		defer conn.Close()

		client := zqbrecord.NewRecordClient(conn)
		err = NotifyHit(client, &zqbrecord.RecordNotifyHitReq{}, 1)
		if err != nil {
			log.Fatalf("printLists.err: %v", err)
		}
		wg.Done()
	}()

	go func() {
		time.Sleep(time.Second * 5)
		connection, err := grpc.Dial(":"+"17802", grpc.WithInsecure())
		if err != nil {
			log.Fatalf("grpc.Dial err: %v", err)
		}

		defer connection.Close()

		client := zqbrecord.NewRecordClient(connection)
		err = NotifyHit(client, &zqbrecord.RecordNotifyHitReq{}, 2)
		if err != nil {
			log.Fatalf("printLists.err: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()
}

func NotifyHit(cli zqbrecord.RecordClient, r *zqbrecord.RecordNotifyHitReq, id int) error{
	stream, err := cli.NotifyHit(context.Background(), r)
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

		log.Printf("id: %v", id)
		log.Printf("resp: recordID: %s, qrcode: %v", resp.RecordId, resp.QrCode)
	}

	return nil
}
