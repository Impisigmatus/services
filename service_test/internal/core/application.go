package core

import "github.com/sirupsen/logrus"

type Application struct{}

func NewApplication() *Application {
	return &Application{}
}

func (app *Application) Ping() {
	logrus.Info("Ping")
}
