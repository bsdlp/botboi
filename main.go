package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/bsdlp/botboi/pkg/bot"
	"github.com/bsdlp/botboi/pkg/cfg"
	"github.com/bsdlp/botboi/pkg/handlers"
	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap"
)

func main() {
	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatal(err)
	}

	var config cfg.Config
	err = envconfig.Process("botboi", &config)
	if err != nil {
		logger.Fatal(err.Error())
	}

	sugaredLogger := logger.Sugar()
	bt := &bot.Bot{
		Config:   config,
		Logger:   sugaredLogger,
		Handlers: &handlers.Loader{Logger: sugaredLogger},
	}

	err = bt.Run()
	if err != nil {
		logger.Fatal(err.Error())
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM)
	<-sc

	err = bt.Stop()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
