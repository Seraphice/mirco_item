package main

import (
	"git.imooc.com/dfs/my_mirco_item/handler"

	"github.com/micro/go-micro/v2/logger"
	"github.com/micro/micro/v2/service"
	//"github.com/micro/micro/v2/service/logger"
)

func main() {
	// Create service
	srv := service.New(
		service.Name("user"),
	)

	// Register handler
	pb.RegisterUserHandler(srv.Server(), handler.New())

	// Run service
	if err := srv.Run(); err != nil {
		logger.Fatal(err)
	}
}
