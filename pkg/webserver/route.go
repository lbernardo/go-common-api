package webserver

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
)

type Route interface {
	Method() string
	Pattern() string
	Run(c *gin.Context)
	Middleware() gin.HandlerFunc
}

func asRoute(f any) any {
	return fx.Annotate(
		f,
		fx.As(new(Route)),
		fx.ResultTags(`group:"routes"`),
	)
}
