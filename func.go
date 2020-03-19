package eglm

import (
	"github.com/GameWith/gwlog"
	"github.com/labstack/echo/v4"
)

// LoggingFunc is write log function
type LoggingFunc func(logger gwlog.Logger, param *Parameter, c echo.Context) error

// DefaultLoggingFunc is default write log function
func DefaultLoggingFunc(logger gwlog.Logger, param *Parameter, _ echo.Context) error {
	p := createAttributeByParameter(param)
	logger.WithFields(p.ToMap()).Info(p.Method + "\x20" + p.Path)
	return nil
}
