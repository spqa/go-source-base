package log

import (
	"go.uber.org/zap"
	"os"
)

var Logger *zap.Logger

func NewLogger() (*zap.Logger, error) {
	env, b := os.LookupEnv("DEBUG")
	if b && env == "true" {
		return zap.NewDevelopment()
	}
	return zap.NewProduction()
}

func init() {
	Logger, _ = NewLogger()
}
