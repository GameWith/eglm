# eglm
gwlog 専用の echo log middleware ライブラリです。

# Usage

echo の Middleware に追加してください。
また、echo の Logger も gwlog.Logger をセットしてください。

```go
e := echo.New()
logger := gwlog.GetLogger()
logger.SetOutput(os.Stdout)
e.Logger = logger
e.Use(eglm.Middleware(&Config{}))
```

# Options

## Skipper

各リクエストで条件に一致した場合ログ出力をスキップします。
デフォルトは**全てスキップしません。**

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

Middleware 経由でログを書き込む処理をします。
ログの書き込み内容をカスタマイズしたい場合に利用します。
デフォルトは `eglm.DefaultLoggingFunc` が設定されています。

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
