package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	tests := []struct {
		handler http.HandlerFunc
		name    string
		route   string
		body    string
	}{
		{
			defaultHandler,
			"Default Handler",
			"/",
			`OK
`,
		},
		{
			statusHandler,
			"Status Handler",
			"/health",
			`{"status": "OK"}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := http.NewRequest("GET", tt.route, nil)
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(tt.handler)

			handler.ServeHTTP(rr, r)

			if status := rr.Code; status != http.StatusOK {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, http.StatusOK)
			}

			expected := tt.body
			if rr.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					rr.Body.String(), expected)
			}
		})
	}
}
