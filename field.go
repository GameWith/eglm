package eglm

import (
	"encoding/json"
)

// DefaultLogField echo default log fields.
type DefaultLogField struct {
	Type         string `json:"type"`
	Status       int    `json:"status"`
	Method       string `json:"method"`
	Path         string `json:"path"`
	UserAgent    string `json:"userAgent"`
	RemoteIP     string `json:"remoteIp"`
	ForwardedFor string `json:"forwardedFor"`
	Latency      int    `json:"latency"`
	Error        string `json:"error"`
	RequestID    string `json:"requestId"`
}

// ToMap struct attributes to map
func (d *DefaultLogField) ToMap() map[string]interface{} {
	return map[string]interface{}{
		"type":         d.Type,
		"method":       d.Method,
		"status":       d.Status,
		"path":         d.Path,
		"userAgent":    d.UserAgent,
		"remoteIp":     d.RemoteIP,
		"forwardedFor": d.ForwardedFor,
		"latency":      d.Latency,
		"error":        d.Error,
		"requestId":    d.RequestID,
	}
}

// ToJSON struct attributes to json
func (d *DefaultLogField) ToJSON() ([]byte, error) {
	return json.Marshal(d)
}
