package eglm

import (
	"reflect"
	"testing"
)

func TestDefaultLogField_ToMap(t *testing.T) {
	type fields struct {
		Type         string
		Status       int
		Method       string
		Path         string
		UserAgent    string
		RemoteIP     string
		ForwardedFor string
		Latency      int
		Error        string
		RequestID    string
	}
	tests := []struct {
		name   string
		fields fields
		want   map[string]interface{}
	}{
		{
			name:   "empty",
			fields: fields{},
			want: map[string]interface{}{
				"type":         "",
				"method":       "",
				"status":       0,
				"path":         "",
				"userAgent":    "",
				"remoteIp":     "",
				"forwardedFor": "",
				"latency":      0,
				"error":        "",
				"requestId":    "",
			},
		},
		{
			name: "all not empty",
			fields: fields{
				Type:         "a",
				Status:       1,
				Method:       "a",
				Path:         "a",
				UserAgent:    "a",
				RemoteIP:     "a",
				ForwardedFor: "a",
				Latency:      1111,
				Error:        "a",
				RequestID:    "a",
			},
			want: map[string]interface{}{
				"type":         "a",
				"method":       "a",
				"status":       1,
				"path":         "a",
				"userAgent":    "a",
				"remoteIp":     "a",
				"forwardedFor": "a",
				"latency":      1111,
				"error":        "a",
				"requestId":    "a",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DefaultLogField{
				Type:         tt.fields.Type,
				Status:       tt.fields.Status,
				Method:       tt.fields.Method,
				Path:         tt.fields.Path,
				UserAgent:    tt.fields.UserAgent,
				RemoteIP:     tt.fields.RemoteIP,
				ForwardedFor: tt.fields.ForwardedFor,
				Latency:      tt.fields.Latency,
				Error:        tt.fields.Error,
				RequestID:    tt.fields.RequestID,
			}
			if got := d.ToMap(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultLogField.ToMap() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDefaultLogField_ToJSON(t *testing.T) {
	type fields struct {
		TraceID      string
		Type         string
		Status       int
		Method       string
		Path         string
		UserAgent    string
		RemoteIP     string
		ForwardedFor string
		Latency      int
		Error        string
		RequestID    string
	}
	tests := []struct {
		name    string
		fields  fields
		want    []byte
		wantErr bool
	}{
		{
			name: "all not empty",
			fields: fields{
				Type:         "a",
				Status:       1,
				Method:       "a",
				Path:         "a",
				UserAgent:    "a",
				RemoteIP:     "a",
				ForwardedFor: "a",
				Latency:      11111,
				Error:        "a",
				RequestID:    "a",
			},
			want:    []byte(`{"type":"a","status":1,"method":"a","path":"a","userAgent":"a","remoteIp":"a","forwardedFor":"a","latency":11111,"error":"a","requestId":"a"}`),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &DefaultLogField{
				Type:         tt.fields.Type,
				Status:       tt.fields.Status,
				Method:       tt.fields.Method,
				Path:         tt.fields.Path,
				UserAgent:    tt.fields.UserAgent,
				RemoteIP:     tt.fields.RemoteIP,
				ForwardedFor: tt.fields.ForwardedFor,
				Latency:      tt.fields.Latency,
				Error:        tt.fields.Error,
				RequestID:    tt.fields.RequestID,
			}
			got, err := d.ToJSON()
			if (err != nil) != tt.wantErr {
				t.Errorf("DefaultLogField.ToJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DefaultLogField.ToJSON() = %v, want %v", string(got), string(tt.want))
			}
		})
	}
}
