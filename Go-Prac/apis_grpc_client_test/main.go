package main

import (
	"context"
	"google.golang.org/grpc"
	"io"
	"log"
	"sync"
	"vincent/practice/Go-Prac/apis_grpc_client_test/zqbrecord"
)

func main() {
	conn, err := grpc.Dial(":"+"17802", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc.Dial err: %v", err)
	}

	defer conn.Close()

	client := zqbrecord.NewRecordClient(conn)


	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		err = NotifyHit(client, &zqbrecord.RecordNotifyHitReq{})
		if err != nil {
			log.Fatalf("printLists.err: %v", err)
		}
		wg.Done()
	}()

	go func() {
		err = NotifyHit(client, &zqbrecord.RecordNotifyHitReq{})
		if err != nil {
			log.Fatalf("printLists.err: %v", err)
		}
		wg.Done()
	}()

	wg.Wait()
}

func NotifyHit(cli zqbrecord.RecordClient, r *zqbrecord.RecordNotifyHitReq) error{
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

		log.Printf("resp: recordID: %s, qrcode: %v", resp.RecordId, resp.QrCode)
	}

	return nil
}
