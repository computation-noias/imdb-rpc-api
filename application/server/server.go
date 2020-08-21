package server

import (
	"context"
	"fmt"
	"github.com/computation-noias/imdb-rpc-api/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

// ServiceInput ...
type ServiceInput struct {
	Port int
}

// Service ...
type Service struct {
	in ServiceInput
	server *grpc.Server
}

// NewService ...
func NewService(in ServiceInput) (*Service, error) {
	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)

	return &Service{
		in: in,
		server: grpcServer,
	}, nil
}

// registerResolvers ...
func (s Service) registerResolvers() {
	pb.RegisterMovieServiceServer(s.server, &MovieResolver{})
}

// Run ...
func (s Service) Run(ctx context.Context) error {
	s.registerResolvers()

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", s.in.Port))
	if err != nil {
		log.Fatalf("Cannot start the GRPC Server: %v", err)
		return err
	}

	err = s.server.Serve(listener)
	if err != nil {
		log.Fatalf("Cannot server the GRPC server: %v", err)
		return err
	}

	return nil
}
