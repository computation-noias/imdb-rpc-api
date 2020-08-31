package main

import (
	"context"
	"fmt"
	"github.com/computation-noias/imdb-rpc-api/application/server"
	"os"
	"strconv"
)

type config struct {
	port int
}

func env() (*config, error) {
	c := &config{}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return nil, err
	}
	c.port = port

	return c, nil
}

func main() {
	ctx := context.Background()

	env, err := env()
	if err != nil {
		fmt.Println("Error during environment variables build", err.Error())
		return
	}

	svr, err := server.NewService(server.ServiceInput{
		Port: env.port,
	})
	if err != nil {
		fmt.Println("Error when creating the GRPC server", err.Error())
		return
	}

	err = svr.Run(ctx)
	if err != nil {
		fmt.Println("Error when starting up the GRPC server", err.Error())
		return
	}

}
