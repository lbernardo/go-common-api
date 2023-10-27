package webserver

import (
	"github.com/gin-contrib/cors"
	"go.uber.org/fx"
)

type WebServer struct {
	routes     []any
	middleware []any
	corsConfig *cors.Config
}

func NewWebServer() *WebServer {
	return &WebServer{
		routes:     []any{},
		middleware: []any{},
	}
}

func (w *WebServer) AddRoute(route ...any) *WebServer {
	for _, r := range route {
		w.routes = append(w.routes, asRoute(r))
	}
	return w
}

func (w *WebServer) RegisterMiddleware(middleare ...any) *WebServer {
	for _, m := range middleare {
		w.middleware = append(w.middleware, m)
	}
	return w
}

func (w *WebServer) SetCors(cfg *cors.Config) {
	w.corsConfig = cfg
}

func (w *WebServer) GetCors() cors.Config {
	if w.corsConfig == nil {
		cfg := cors.DefaultConfig()
		cfg.AllowAllOrigins = true
		cfg.AllowMethods = []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"}
		cfg.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization", "Public-Request", "Accept", "Token"}
		return cfg
	}
	return *w.corsConfig
}

func (w *WebServer) StartModule() fx.Option {
	return fx.Module("webserver",
		fx.Provide(fx.Annotate(newEngine, fx.ParamTags(`group:"routes"`, `name:"cors"`))),
		fx.Provide(fx.Annotate(
			w.GetCors,
			fx.ResultTags(`name:"cors"`),
		)),
		fx.Provide(w.middleware...),
		fx.Provide(w.GetCors),
		fx.Provide(w.routes...),
		fx.Invoke(newServer))
}
