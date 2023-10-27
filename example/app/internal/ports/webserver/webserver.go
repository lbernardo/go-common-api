package webserver

import (
	"github.com/lbernardo/go-common-api/example/app/internal/ports/webserver/handlers"
	"go.uber.org/fx"
)

func NewWebServer() fx.Option {
	w := webserver.NewWebServer()
	w.AddRoute(handlers.NewExample, handlers.NewExample2)
	return w.StartModule()
}
