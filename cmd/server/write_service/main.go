package main

import (
	"cqrs-grpc-test/config"
	"cqrs-grpc-test/internal/server"
	"cqrs-grpc-test/pkg/logger"
	"flag"
	"log"
)

func main() {
	flag.Parse()
	cfg, err := config.InitConfig()
	if err != nil {
		log.Fatal(err)
	}
	appLogger := logger.NewAppLogger(cfg.Logger)
	appLogger.InitLogger()
	appLogger.WithName(cfg.ServiceName)

	s := server.NewServer(appLogger, cfg)
	appLogger.Fatal(s.Run("write"))

}
