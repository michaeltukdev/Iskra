package main

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
)

var routes = []string{"/api/v1/meminfo", "/api/v1/cpuinfo", "/api/v1/uptime"}

func TestEndpoints(t *testing.T) {
	e := echo.New()
	registerRoutes(e)

	for _, route := range routes {
		req := httptest.NewRequest("GET", route, nil)
		rec := httptest.NewRecorder()
		e.ServeHTTP(rec, req)

		if rec.Code != 200 {
			t.Fatalf("Failed to fetch %s received error code %d", route, rec.Code)
		}
	}
}
