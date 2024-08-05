package main

import (
	"client/internal/api/client"
	readerstream "client/internal/api/readerStream"
	"client/internal/config"
	"client/internal/services"
	"client/internal/storage"
	"client/internal/utils"
	"client/storage/postgres"
	"context"
	"flag"
	"log"
	"time"
)

func main() {

	var k int
	flag.IntVar(&k, "k", 3, "- an STD anomaly coefficient")
	flag.Parse()

	ctx, _ := context.WithDeadline(context.Background(), time.Now().Add(time.Second*100))

	cfg := config.MustReadConfig()

	db := postgres.MustConnect(cfg)

	strg := storage.NewFrequency(db)

	err := strg.CreateTable(context.Background())

	if err != nil {
		log.Fatal("not create table", err)
	}

	stat := utils.NewStat()

	service := services.NewService(strg, stat, k)

	cl, err := client.ConnectClient(ctx, service, cfg)

	if err != nil {
		log.Fatal("not connect to the grpc server", err)
	}

	defer cl.Close()

	doneTicker := make(chan bool)
	utils.StartTicker(doneTicker, stat)
	readerstream.ReadStream(cl, doneTicker, &utils.DefaultErrorHandler{})
}
