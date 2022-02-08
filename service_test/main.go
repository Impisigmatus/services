package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/Impisigmatus/services/service_test/internal/config"
	"github.com/Impisigmatus/services/service_test/internal/core"
	"github.com/sirupsen/logrus"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		logrus.Panic("Invalid config: ", err)
	}

	app := core.NewApplication()
	service := core.NewService(app, cfg.Port)

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		service.Stop()
	}()

	if err := service.Start(); err != nil {
		logrus.Fatalf("Service error: %s", err)
	}
}
