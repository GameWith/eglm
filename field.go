package eglm

import (
	"encoding/json"
)

// AccessLogField echo default log fields.
type AccessLogField struct {
	Type         string `json:"type"`
	Status       int    `json:"status"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	UserAgent    string `json:"userAgent"`
	RemoteIP     string `json:"remoteIp"`
	ForwardedFor string `json:"forwardedFor"`
	Elapsed      int    `json:"elapsed"`
	Error        string `json:"error"`
	RequestID    string `json:"requestId"`
}

// ToMap struct attributes to map
func (a *AccessLogField) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"type":         a.Type,
		"status":       a.Status,
		"method":       a.Method,
		"path":         a.Path,
		"userAgent":    a.UserAgent,
		"remoteIp":     a.RemoteIP,
		"forwardedFor": a.ForwardedFor,
		"elapsed":      a.Elapsed,
		"error":        a.Error,
		"requestId":    a.RequestID,
	}
}

// ToJSON struct attributes to json
func (a *AccessLogField) ToJSON() ([]byte, error) {
	return json.Marshal(a)
}
