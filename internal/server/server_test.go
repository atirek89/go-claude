package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHelloHandler(t *testing.T) {
	mux := New()

	tests := []struct {
		name     string
		url      string
		wantBody string
		wantCode int
	}{
		{name: "no name param", url: "/hello", wantBody: "Hello, World!", wantCode: http.StatusOK},
		{name: "name param set", url: "/hello?name=Atirek", wantBody: "Hello, Atirek!", wantCode: http.StatusOK},
		{name: "whitespace name falls back", url: "/hello?name=+", wantBody: "Hello, World!", wantCode: http.StatusOK},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(http.MethodGet, tt.url, nil)
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)

			if rec.Code != tt.wantCode {
				t.Errorf("status = %d, want %d", rec.Code, tt.wantCode)
			}
			if body := strings.TrimSpace(rec.Body.String()); body != tt.wantBody {
				t.Errorf("body = %q, want %q", body, tt.wantBody)
			}
		})
	}
}
