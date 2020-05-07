package eglm

// Parameter is LoggingFunc parameter struct
type Parameter struct {
	ID        string
	Method    string
	Host      string
	URI       string
	Status    int
	Path      string
	RemoteIP  string
	Referer   string
	UserAgent string
	Elapsed   int
	BytesIn   string
	BytesOut  string
	Error     error
}

func ConvertAccessLogFieldByParameter(param *Parameter) *AccessLogField {
	f := &AccessLogField{
		Type:      "ACCESS",
		Status:    param.Status,
		Method:    param.Method,
		Path:      param.Path,
		UserAgent: param.UserAgent,
		RemoteIP:  param.RemoteIP,
		Elapsed:   param.Elapsed,
		RequestID: param.ID,
	}
	if param.Error != nil {
		f.Error = param.Error.Error()
	}
	return f
}
