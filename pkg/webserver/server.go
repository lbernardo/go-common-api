package webserver

import (
	"context"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
	"os"
)

func newServer(lc fx.Lifecycle, engine *gin.Engine, logger *zap.Logger, corsConfig cors.Config) {
	port := os.Getenv("PORT")
	if port == "" {
		port = "80"
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			logger.Info(fmt.Sprintf("start webserver on port :%v", port))

			logger.Info("CORS setup", zap.Strings("AllowOrigins", corsConfig.AllowOrigins),
				zap.Bool("AllowAllOrigins", corsConfig.AllowAllOrigins),
				zap.Strings("AllowMethods", corsConfig.AllowMethods),
				zap.Strings("AllowHeaders", corsConfig.AllowHeaders))
			engine.Use(cors.New(corsConfig))
			go func() {
				if err := engine.Run(fmt.Sprintf("0.0.0.0:%v", port)); err != nil {
					logger.Error("error to start webserver", zap.Error(err))
					os.Exit(1)
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			logger.Info("stopping webserver")
			return nil
		},
	})
}
