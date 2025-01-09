package main

import (
	note "app/pkg/v1"
	"context"
	"log"
	"time"

	"github.com/fatih/color"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	NoteID  = 12
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to connect to server: %v", err)
	}
	defer conn.Close()

	c := note.NewV1Client(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*100)
	defer cancel()
	count := 0
	for {
		r, err := c.Get(ctx, &note.GetRequest{Id: NoteID})
		if err != nil {
			log.Fatalf("failed to get note by id: %v", err)
		}
		log.Printf(color.RedString("Note info:\n"), color.GreenString("%+v", r.GetNote()))
		count++
		if count == 10000000 {
			break
		}
	}

}
