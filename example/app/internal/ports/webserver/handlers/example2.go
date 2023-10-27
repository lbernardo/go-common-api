package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Example2 struct {
}

func NewExample2() *Example2 {
	return &Example2{}
}

func (e *Example2) Method() string {
	return http.MethodGet
}

func (e *Example2) Pattern() string {
	return "/example2"
}

func (e *Example2) Run(c *gin.Context) {
	c.Status(http.StatusOK)
}

func (e *Example2) Middleware() gin.HandlerFunc {
	return nil
}
