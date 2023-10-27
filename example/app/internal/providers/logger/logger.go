package logger

import "go.uber.org/zap"

func NewLogger() *zap.Logger {
	l, _ := zap.NewDevelopment()
	return l
}
