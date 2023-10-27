package webserver

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/lbernardo/go-common-api/pkg/models"
	"net/http"
)

func newEngine(routes []Route, corsConfig cors.Config) *gin.Engine {

	gin.SetMode(gin.ReleaseMode)
	engine := gin.Default()
	for _, route := range routes {
		if route.Middleware() == nil {
			g := engine.Group("/")
			g.Use(cors.New(corsConfig))
			g.Handle(route.Method(), route.Pattern(), route.Run)
			continue
		}
		g := engine.Group("/")
		g.Use(route.Middleware())
		g.Use(cors.New(corsConfig))
		g.Handle(route.Method(), route.Pattern(), route.Run)
	}
	engine.GET("/-/service/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, models.NewHealth())
	})
	return engine
}
