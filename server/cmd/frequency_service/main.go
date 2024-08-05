package main

import (
	grpcfrequency "frequency_service/internal/app/grpc/frequency"
	"frequency_service/internal/config"
	"frequency_service/internal/services"
	"log"
	"log/slog"
	"net"
	"os"

	"google.golang.org/grpc"
)

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	cfg := config.MustReadConfig()
	listener, err := net.Listen("tcp", cfg.Host+cfg.Port)

	if err != nil {
		log.Fatal("error building server" + err.Error())
	}
	s := grpc.NewServer()
	frequencyService := services.NewServices()
	grpcfrequency.Register(s, frequencyService)

	log.Println("start server")

	if err := s.Serve(listener); err != nil {
		log.Fatal("error building server" + err.Error())
	}

}
