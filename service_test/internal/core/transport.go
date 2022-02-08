package core

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	app *Application
}

func NewTransport(app *Application) *Transport {
	return &Transport{app: app}
}

func (transport *Transport) Ping(ctx *gin.Context) {
	transport.app.Ping()
	ctx.String(http.StatusNoContent, "")
}
