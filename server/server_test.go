package server

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

type Test struct {
	name          string
	server        *httptest.Server
	response      *Weather
	expectedError error
}

func TestGetWeather(t *testing.T) {
	tests := []Test{
		{
			name: "basic-request",
			server: httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				w.WriteHeader(http.StatusOK)
				w.Write([]byte(`{ "city": "London, UK", "forecast": "gloomy" }`))
			})),
			response: &Weather{
				City:     "London, UK",
				Forecast: "gloomy",
			},
			expectedError: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			defer test.server.Close()

			resp, err := GetWeather(test.server.URL)

			if !reflect.DeepEqual(resp, test.response) {
				t.Errorf("Expected %v. Got %v", test.response, resp)
			}

			if !errors.Is(err, test.expectedError) {
				t.Errorf("Expected error %v. Got %v", test.expectedError, err)
			}
		})
	}
}
