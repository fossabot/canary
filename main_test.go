package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestHandlers(t *testing.T) {
	tt := []struct {
		name    string
		handler http.HandlerFunc
		route   string
		body    string
	}{
		{
			"Default Handler",
			defaultHandler,
			"/",
			"OK\n",
		},
		{
			"Status Handler",
			statusHandler,
			"/health",
			"{\"status\": \"OK\"}\n",
		},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			req, err := http.NewRequest("GET", tc.route, nil)
			if err != nil {
				t.Fatal(err)
			}

			res := httptest.NewRecorder()
			handler := http.HandlerFunc(tc.handler)

			handler.ServeHTTP(res, req)

			if status := res.Code; status != http.StatusOK {
				t.Errorf("handler returned status code: got %v want %v",
					status, http.StatusOK)
			}

			expected := tc.body
			if res.Body.String() != expected {
				t.Errorf("handler returned unexpected body: got %v want %v",
					res.Body.String(), expected)
			}
		})
	}
}
