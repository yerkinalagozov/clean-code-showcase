package main

import (
	"context"
	"log"
	"time"

	"github.com/yerkinalagozov/clean-code-showcase.git/configs"
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/dbconnection"
	handler2 "github.com/yerkinalagozov/clean-code-showcase.git/internal/user/controller/handler"
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/infra/repository"
	"github.com/yerkinalagozov/clean-code-showcase.git/internal/user/service"
)

const (
	defaultTimeout = 5 * time.Second
)

func main() {
	cfg, err := configs.NewConfig()
	if err != nil {
		log.Println("failed to create config: ", err)
		return
	}
	ctx, cancel := context.WithTimeout(context.Background(), defaultTimeout)
	defer cancel()
	pool, err := dbconnection.NewConnection(ctx, cfg)
	if err != nil {
		log.Println("failed to create connection: ", err)
		return
	}
	defer pool.Close()

	orderRepoImpl := repository.NewOrderRepo(pool)
	userRepoImpl := repository.NewUserRepo(pool)

	productImpl := repository.NewProductRepo(pool)
	orderItemsImpl := repository.NewOrderItemsRepo(pool)

	serv := service.NewService(userRepoImpl, productImpl, orderRepoImpl, orderItemsImpl)

	handler := handler2.NewHandler(serv)

	handler.StartServer(cfg)
}
