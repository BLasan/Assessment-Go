package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/BLasan/Assessment-Go/handlers"
)

func TestHelloWorldHandler(t *testing.T) {
	tests := []struct {
		name           string
		queryParam     string
		expectedStatus int
		expectedBody   interface{}
	}{
		// Test empty parameter
		{
			name:           "Empty name parameter",
			queryParam:     "",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   handlers.ErrorResponse{Error: "Invalid Input"},
		},
		// Test invalid characters (numbers)
		{
			name:           "Name with numbers",
			queryParam:     "John123",
			expectedStatus: http.StatusOK,
			expectedBody:   handlers.MessageResponse{Message: "Hello John123"},
		},
		// Test invalid characters (special characters)
		{
			name:           "Name with special characters",
			queryParam:     "John@Doe",
			expectedStatus: http.StatusOK,
			expectedBody:   handlers.MessageResponse{Message: "Hello John@Doe"},
		},
		// Test invalid characters (spaces) in the middle
		{
			name:           "Name with spaces",
			queryParam:     "John Doe",
			expectedStatus: http.StatusOK,
			expectedBody:   handlers.MessageResponse{Message: "Hello John Doe"},
		},
		// Test names starting with A-M (lowercase)
		{
			name:           "Name starting with 'a' (lowercase)",
			queryParam:     "alice",
			expectedStatus: http.StatusOK,
			expectedBody:   handlers.MessageResponse{Message: "Hello alice"},
		},
		{
			name:           "Name starting with 'A' (uppercase)",
			queryParam:     "Alice",
			expectedStatus: http.StatusOK,
			expectedBody:   handlers.MessageResponse{Message: "Hello Alice"},
		},
		{
			name:           "Name starting with 'm' (lowercase)",
			queryParam:     "mary",
			expectedStatus: http.StatusOK,
			expectedBody:   handlers.MessageResponse{Message: "Hello mary"},
		},
		{
			name:           "Name starting with 'M' (uppercase)",
			queryParam:     "Mary",
			expectedStatus: http.StatusOK,
			expectedBody:   handlers.MessageResponse{Message: "Hello Mary"},
		},
		{
			name:           "Name starting with 'b'",
			queryParam:     "bob",
			expectedStatus: http.StatusOK,
			expectedBody:   handlers.MessageResponse{Message: "Hello bob"},
		},
		// Test names starting with N-Z (lowercase)
		{
			name:           "Name starting with 'n' (lowercase)",
			queryParam:     "nancy",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   handlers.ErrorResponse{Error: "Invalid Input"},
		},
		{
			name:           "Name starting with 'N' (uppercase)",
			queryParam:     "Nancy",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   handlers.ErrorResponse{Error: "Invalid Input"},
		},
		{
			name:           "Name starting with 'z' (lowercase)",
			queryParam:     "zoe",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   handlers.ErrorResponse{Error: "Invalid Input"},
		},
		{
			name:           "Name starting with 'Z' (uppercase)",
			queryParam:     "Zoe",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   handlers.ErrorResponse{Error: "Invalid Input"},
		},
		{
			name:           "Name starting with 'p'",
			queryParam:     "peter",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   handlers.ErrorResponse{Error: "Invalid Input"},
		},
		// Test names starting with numbers
		{
			name:           "Name starting with '1'",
			queryParam:     "1peter",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   handlers.ErrorResponse{Error: "Invalid Input"},
		},
		// Test names starting with non alphabetic characters
		{
			name:           "Name starting with '1'",
			queryParam:     "*peter",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   handlers.ErrorResponse{Error: "Invalid Input"},
		},
		// Test names starting trailing/leading spaces
		{
			name:           "Name starting with leading spaces",
			queryParam:     " aeter",
			expectedStatus: http.StatusOK,
			expectedBody:   handlers.MessageResponse{Message: "Hello aeter"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create a request
			req, err := http.NewRequest("GET", "/hello-world?name="+tt.queryParam, nil)
			if err != nil {
				t.Fatal(err)
			}

			// Create a ResponseRecorder
			rr := httptest.NewRecorder()

			// Call the handler
			handler := http.HandlerFunc(handlers.HelloWorldHandler)
			handler.ServeHTTP(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v",
					status, tt.expectedStatus)
			}

			// Check response body
			if tt.expectedStatus == http.StatusOK {
				var response handlers.MessageResponse
				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}
				expected := tt.expectedBody.(handlers.MessageResponse)
				if response.Message != expected.Message {
					t.Errorf("handler returned unexpected body: got %v want %v",
						response.Message, expected.Message)
				}
			} else {
				var response handlers.ErrorResponse
				if err := json.NewDecoder(rr.Body).Decode(&response); err != nil {
					t.Errorf("Failed to decode response: %v", err)
				}
				expected := tt.expectedBody.(handlers.ErrorResponse)
				if response.Error != expected.Error {
					t.Errorf("handler returned unexpected body: got %v want %v",
						response.Error, expected.Error)
				}
			}
		})
	}
}
