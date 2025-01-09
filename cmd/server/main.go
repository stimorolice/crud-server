package main

import (
	note "app/pkg/v1"
	"context"
	"log"
	"net"

	"github.com/brianvoe/gofakeit"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/types/known/timestamppb"
)

const port = ":50051"

type Server struct {
	note.UnimplementedV1Server
}

func (s *Server) Get(ctx context.Context, req *note.GetRequest) (*note.GetResponse, error) {
	log.Printf("Note id: %d", req.GetId())

	return &note.GetResponse{
		Note: &note.Note{
			Id: req.GetId(),
			Info: &note.NoteInfo{
				Title:    gofakeit.BeerMalt(),
				Content:  gofakeit.BS(),
				Author:   gofakeit.Email(),
				IsPublic: gofakeit.Bool(),
			},
			CreatedAt: timestamppb.Now(),
			UpdatedAt: timestamppb.New(gofakeit.Date()),
		},
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	note.RegisterV1Server(s, &Server{})
	log.Println("started service bro... on port", port)
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve shit bruh: %v", err)
	}

}
