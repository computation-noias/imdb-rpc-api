package server

import (
	"context"
	"fmt"
	"github.com/computation-noias/imdb-rpc-api/pb"
	"github.com/google/uuid"
)

// MovieResolver ...
type MovieResolver struct {}

func (m MovieResolver) Create(ctx context.Context, req *pb.CreateMoviePayload) (*pb.CreateMovieResponse, error) {
	fmt.Println("Received a request to create a Movie.")
	fmt.Printf("Movie name: %s.\n", req.Data.Name)
	fmt.Printf("Movie synopsis: %s.\n", req.Data.Synopsis)

	return &pb.CreateMovieResponse{
		MovieId: uuid.New().String(),
	}, nil
}