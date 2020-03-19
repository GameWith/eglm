package eglm

import (
	em "github.com/labstack/echo/v4/middleware"
)

// Config is Middleware config
type Config struct {
	Skipper     em.Skipper
	LoggingFunc LoggingFunc
}
