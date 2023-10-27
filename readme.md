# common-webserver

Common webserver using uber fx (https://github.com/uber-go/fx)

## install

```shell
go get github.com/lbernardo/go-common-api
```

## Usage

`cmd/webserver/main.go`
```go
package main

import (
	"github.com/lbernardo/go-common-api/example/app/internal/ports/webserver"
	"github.com/lbernardo/go-common-api/example/app/internal/providers/logger"
	"go.uber.org/fx"
)

func main() {
	app := fx.New(
		fx.Provide(logger.NewLogger), // is need
		webserver.NewWebServer(), // 
	)
	app.Run()
}
```

`internal/ports/webserver/server.go`

```go
package webserver

import (
	"github.com/lbernardo/go-common-api/example/app/internal/ports/webserver/handlers"
	"github.com/lbernardo/go-common-api/pkg/webserver"
	"go.uber.org/fx"
)

func NewWebServer() fx.Option {
	w := webserver.NewWebServer()
	// add route
	w.AddRoute(handlers.NewExample)
	// return fx module
	return w.StartModule()
}
```

`interal/ports/webserver/handlers/example.go`

```go
package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Example struct {
}


func NewExample() *Example {
	return &Example{}
}

// Method
func (e *Example) Method() string {
	return http.MethodGet
}

// Pattern url
func (e *Example) Pattern() string {
	return "/example"
}

// Use middleware to authorization or loggs
func (e *Example) Middleware() gin.HandlerFunc {
	return nil
}

// Run
func (e *Example) Run(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]int64{"unix": time.Now().Unix()})
}
```

## Configure CORS

By default, cors is Allow All, but if you need configure, you can use https://github.com/gin-contrib/cors

```go
package webserver

import (
	"github.com/gin-contrib/cors"
	"github.com/lbernardo/go-common-api/example/app/internal/ports/webserver/handlers"
	"github.com/lbernardo/go-common-api/pkg/webserver"
	"go.uber.org/fx"
)

func NewWebServer() fx.Option {
	w := webserver.NewWebServer()
	// add route
	w.AddRoute(handlers.NewExample)
	w.SetCors(&cors.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
    })
	// return fx module
	return w.StartModule()
}
```