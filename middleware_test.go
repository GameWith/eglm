package eglm

import (
	"bytes"
	"fmt"
	"github.com/GameWith/gwlog/formatter"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/GameWith/gwlog"
	"github.com/labstack/echo/v4"
)

func TestMiddleware(t *testing.T) {
	e := echo.New()
	buf := new(bytes.Buffer)
	logger := gwlog.GetLogger()
	logger.SetOutput(buf)
	e.Logger = logger
	e.Use(Middleware(&Config{}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	status, _ := request(http.MethodGet, "/", e)
	if status != http.StatusOK {
		t.Errorf("got = %v, want = %v", status, http.StatusOK)
	}
	if strings.Contains(buf.String(), "type=ACCESS") == false {
		t.Errorf("invalid buf: %v", buf.String())
	}
}

func TestMiddleware_skipper(t *testing.T) {
	e := echo.New()
	buf := new(bytes.Buffer)
	logger := gwlog.GetLogger()
	logger.SetOutput(buf)
	e.Logger = logger
	e.Use(Middleware(&Config{
		Skipper: func(c echo.Context) bool {
			return c.Path() == "/"
		},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	status, _ := request(http.MethodGet, "/", e)
	if status != http.StatusOK {
		t.Errorf("got = %v, want = %v", status, http.StatusOK)
	}
	if 0 != buf.Len() {
		t.Errorf("invalid buf: %v", buf.String())
	}
}

func TestMiddleware_loggingFunc(t *testing.T) {
	e := echo.New()
	buf := new(bytes.Buffer)
	logger := gwlog.GetLogger()
	logger.SetOutput(buf)
	e.Logger = logger
	e.Use(Middleware(&Config{
		LoggingFunc: func(logger gwlog.Logger, _ *Parameter, _ echo.Context) error {
			logger.Print("a")
			return nil
		},
	}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	status, _ := request(http.MethodGet, "/", e)
	if status != http.StatusOK {
		t.Errorf("got = %v, want = %v", status, http.StatusOK)
	}
	if strings.Contains(buf.String(), "msg=a") == false {
		t.Errorf("invalid buf: %v", buf.String())
	}
}

func TestMiddleware_error(t *testing.T) {
	e := echo.New()
	buf := new(bytes.Buffer)
	logger := gwlog.GetLogger()
	logger.SetOutput(buf)
	e.Logger = logger
	e.Use(Middleware(&Config{}))
	e.GET("/", func(c echo.Context) error {
		return fmt.Errorf("ERROR")
	})
	status, _ := request(http.MethodGet, "/", e)
	if status != http.StatusInternalServerError {
		t.Errorf("got = %v, want = %v", status, http.StatusInternalServerError)
	}
	if strings.Contains(buf.String(), "error=ERROR") == false {
		t.Errorf("NotFound Text: ERROR")
	}
}

func TestMiddleware_json(t *testing.T) {
	e := echo.New()
	buf := new(bytes.Buffer)
	logger := gwlog.GetLogger()
	logger.SetFormatter(&formatter.JSONFormatter{})
	logger.SetOutput(buf)
	e.Logger = logger
	e.Use(Middleware(&Config{}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	status, _ := request(http.MethodGet, "/", e)
	if status != http.StatusOK {
		t.Errorf("got = %v, want = %v", status, http.StatusOK)
	}
	if strings.Contains(buf.String(), `"message":"GET /"`) == false {
		t.Errorf("NotFound Text: ERROR")
	}
}

func TestMiddleware_defaultLogger(t *testing.T) {
	e := echo.New()
	e.Use(Middleware(&Config{}))
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})
	status, body := request(http.MethodGet, "/", e)
	if status != http.StatusInternalServerError {
		t.Errorf("got = %v, want = %v", status, http.StatusInternalServerError)
	}
	bw := `{"message":"invalid instance type. only gwlog.Logger"}` + "\n"
	if body != bw {
		t.Errorf("got = %v, want = %v", body, bw)
	}
}

func request(method, path string, e *echo.Echo) (int, string) {
	req := httptest.NewRequest(method, path, nil)
	rec := httptest.NewRecorder()
	e.ServeHTTP(rec, req)
	return rec.Code, rec.Body.String()
}
