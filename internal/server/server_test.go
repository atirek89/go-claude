package server

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestGoodbyeHandler(t *testing.T) {
	mux := New()

	tests := []struct {
		name     string
		url      string
		wantBody string
		wantCode int
	}{
		{name: "no name param", url: "/goodbye", wantBody: "Goodbye, World!", wantCode: http.StatusOK},
		{name: "name param set", url: "/goodbye?name=Atirek", wantBody: "Goodbye, Atirek!", wantCode: http.StatusOK},
		{name: "whitespace name falls back", url: "/goodbye?name=+", wantBody: "Goodbye, World!", wantCode: http.StatusOK},
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

func TestHealthHandler(t *testing.T) {
	mux := New()
	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusOK {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusOK)
	}
	if ct := rec.Header().Get("Content-Type"); ct != "application/json" {
		t.Errorf("Content-Type = %q, want %q", ct, "application/json")
	}
	if body := strings.TrimSpace(rec.Body.String()); body != `{"status":"ok"}` {
		t.Errorf("body = %q, want %q", body, `{"status":"ok"}`)
	}
}

func TestNotFoundHandler(t *testing.T) {
	mux := New()
	req := httptest.NewRequest(http.MethodGet, "/unknown", nil)
	rec := httptest.NewRecorder()
	mux.ServeHTTP(rec, req)

	if rec.Code != http.StatusNotFound {
		t.Errorf("status = %d, want %d", rec.Code, http.StatusNotFound)
	}
}

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
