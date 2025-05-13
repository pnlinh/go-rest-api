package main

import (
	"bytes"
	"go-rest-api/pkg/recipes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainEndpoints(t *testing.T) {
	tests := []struct {
		name       string
		url        string
		method     string
		body       string
		statusCode int
	}{
		{name: "livez ok", url: "/livez", method: "GET", statusCode: http.StatusNoContent},
		{name: "readyz ok", url: "/readyz", method: "GET", statusCode: http.StatusNoContent},
		{name: "list recipes ok", url: "/recipes", method: "GET", statusCode: http.StatusOK},
		{name: "create recipe ok", url: "/recipes", method: "POST", body: "{\"name\":\"Spaghetti Carbonara\",\"ingredients\":[{\"name\":\"tomatoes\"},{\"name\":\"onion\"},{\"name\":\"garlic\"},{\"name\":\"oil\"}]}", statusCode: http.StatusCreated},
	}

	store := recipes.NewMemStore()
	recipesHandler := NewRecipesHandler(store)
	mux := http.NewServeMux()

	mux.Handle("/livez", &healthCheckHandler{})
	mux.Handle("/readyz", &healthCheckHandler{})
	mux.Handle("/recipes", recipesHandler)
	mux.Handle("/recipes/", recipesHandler)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest(tt.method, tt.url, bytes.NewBufferString(tt.body))
			rec := httptest.NewRecorder()
			mux.ServeHTTP(rec, req)

			if rec.Code != tt.statusCode {
				t.Errorf("expected status %d, got %d", tt.statusCode, rec.Code)
			}
		})
	}
}
