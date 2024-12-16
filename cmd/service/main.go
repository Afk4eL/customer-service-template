package main

import (
	"customer-service/config"
	"customer-service/internal/server"
	"customer-service/pkg/logger"
	"customer-service/pkg/postgres"
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	args := os.Args[1:]

	if len(args) < 1 {
		fmt.Println("Usage go run <path to main.go> [arguments] \n Required arguments: \n - Name config file in ./config dir")
		os.Exit(1)
	}

	cfg, err := config.GetConfig(args[0])
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(cfg)

	logger.InitLogger(cfg.Mode)

	pgDb, err := postgres.NewPostgresDb(cfg)
	if err != nil {
		logger.Logger.Fatal().Str("Postgres init failed", err.Error())
		os.Exit(1)
	}

	authServer := server.NewCustomerServer(cfg, pgDb)

	go authServer.Run()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, syscall.SIGTERM, syscall.SIGINT)

	<-stop

	authServer.Stop()
}
