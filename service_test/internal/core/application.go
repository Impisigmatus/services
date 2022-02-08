package core

import (
	"github.com/Impisigmatus/services/service_test/docs"
	"github.com/sirupsen/logrus"
)

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Ping() {
	logrus.Info("Ping")
}

func (app *Application) Info() []byte {
	return docs.README
}
