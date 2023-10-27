package handlers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type Example struct {
}

func NewExample() *Example {
	return &Example{}
}

func (e *Example) Method() string {
	return http.MethodGet
}

func (e *Example) Pattern() string {
	return "/example"
}

func (e *Example) Middleware() gin.HandlerFunc {
	return func(context *gin.Context) {
		fmt.Println("Call middleware")
		context.Next()
	}
}

func (e *Example) Run(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]int64{"unix": time.Now().Unix()})
}
