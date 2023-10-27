package main

import (
	"github.com/lbernardo/go-common-api/example/app/internal/ports/webserver"
	"github.com/lbernardo/go-common-api/example/app/internal/providers/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		//fx.NopLogger,
		fx.Provide(logger.NewLogger),
		webserver.NewWebServer(),
	)
	app.Run()
}
