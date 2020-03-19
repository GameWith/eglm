# eglm

echo log middleware library dedicated to gwlog.

# Usage

Please add it to Middleware of echo.

Also, set gwlog.Logger for the Logger of the echo.

```go
e := echo.New()
logger := gwlog.GetLogger()
logger.SetOutput(os.Stdout)
e.Logger = logger
e.Use(eglm.Middleware(&Config{}))
```

# Options

## Skipper

Log output is skip if the conditions are met in request.

**The default is not to skip everything.**

```go
e := echo.New()
e.Logger = gwlog.GetLogger()
e.Use(eglm.Logger(&Config{
    Skipper: func(c echo.Context) bool {
        return c.Path() == "/"
    },
}))
```

## LoggingFunc

Write the log via Middleware.

Use this when you want to customize the log contents.

If nothing is set, `eglm.DefaultLoggingFunc` will be set.

```go
e := echo.New()
e.Logger = gwlog.GetLogger()
e.Use(eglm.Logger(&Config{
    LoggingFunc: func(logger gwlog.Logger, p *Parameter, c echo.Context) error {
    	logger.Print("a")
        return nil
    },
}))
```
