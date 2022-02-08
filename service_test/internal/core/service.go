package core

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type Service struct {
	server *http.Server
}

func NewService(app *Application, port uint64) *Service {
	mux := gin.New()
	mux.Use(func(ctx *gin.Context) {
		ctx.Writer.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate, max-age=0")

		start := time.Now()
		ctx.Next()
		duration := time.Since(start)

		logrus.Infof("%d | (%s) | %s | %s %s", ctx.Writer.Status(), duration.String(), ctx.ClientIP(), ctx.Request.Method, ctx.Request.URL.Path)
	}, gin.Recovery())

	transport := NewTransport(app)
	mux.GET("/api/ping", transport.Ping)
	mux.GET("/api/info", transport.Info)

	return &Service{
		server: &http.Server{
			Addr:    fmt.Sprintf(":%v", port),
			Handler: mux,
		},
	}
}

func (app *Service) Start() error {
	logrus.Debug("Service started")
	if err := app.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (app *Service) Stop() {
	logrus.Debug("Service stoped")
	app.server.Close()
}
