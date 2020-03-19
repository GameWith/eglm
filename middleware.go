package eglm

import (
	"github.com/GameWith/gwlog"
	"github.com/labstack/echo/v4"
	"strconv"
	"time"
)

// Middleware echo write log middleware
func Middleware(config *Config) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if config != nil && config.Skipper != nil {
				if config.Skipper(c) {
					return next(c)
				}
			}
			fn := DefaultLoggingFunc
			if config != nil && config.LoggingFunc != nil {
				fn = config.LoggingFunc
			}
			return loggingHandle(c, next, fn)
		}
	}
}

func loggingHandle(c echo.Context, next echo.HandlerFunc, fn LoggingFunc) error {
	logger, ok := c.Logger().(gwlog.Logger)
	if !ok {
		err := echo.NewHTTPError(500, "invalid instance type. only gwlog.Logger")
		return err
	}

	req := c.Request()
	res := c.Response()
	start := time.Now()

	var err error
	if err = next(c); err != nil {
		c.Error(err)
	}

	stop := time.Now()

	p := req.URL.Path
	if p == "" {
		p = "/"
	}

	bytesIn := req.Header.Get(echo.HeaderContentLength)

	if bytesIn == "" {
		bytesIn = "0"
	}

	// request id
	id := req.Header.Get(echo.HeaderXRequestID)
	if id == "" {
		id = res.Header().Get(echo.HeaderXRequestID)
	}

	param := &Parameter{
		ID:        id,
		Method:    req.Method,
		Host:      req.Host,
		URI:       req.RequestURI,
		Status:    res.Status,
		Path:      p,
		RemoteIP:  c.RealIP(),
		Referer:   req.Referer(),
		UserAgent: req.UserAgent(),
		Latency:   int(stop.Sub(start).Milliseconds()),
		BytesIn:   bytesIn,
		BytesOut:  strconv.FormatInt(res.Size, 10),
		Error:     err,
	}

	return fn(logger, param, c)
}
