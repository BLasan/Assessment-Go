package handlers

import (
	"encoding/json"
	"net/http"
	"strings"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type MessageResponse struct {
	Message string `json:"message"`
}

// HelloWorldHandler handles GET /hello-world?name=...
func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Only allow GET requests
	if r.Method != http.MethodGet {
		w.Header().Set("Allow", http.MethodGet)
		w.WriteHeader(http.StatusMethodNotAllowed)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Method Not Allowed"})
		return
	}

	// Get the name query parameter
	name := r.URL.Query().Get("name")

	// Check if name is empty
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid Input"})
		return
	}

	// Trim leading/trailing whitespaces
	name = strings.TrimSpace(name)
	if name == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid Input"})
		return
	}

	// Validate only the first character is an English letter (A-Z, a-z)
	firstChar := rune(name[0])
	if (firstChar < 'A' || firstChar > 'Z') && (firstChar < 'a' || firstChar > 'z') {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid Input"})
		return
	}

	// Normalize first character for range check
	firstChar = rune(strings.ToUpper(string(firstChar))[0])

	// Check if first character is between A-M
	if firstChar >= 'A' && firstChar <= 'M' {
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(MessageResponse{Message: "Hello " + name})
		return
	}

	// If first character is between N-Z, return error
	if firstChar >= 'N' && firstChar <= 'Z' {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid Input"})
		return
	}

	// Fallback for any other case (shouldn't happen with our validation)
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(ErrorResponse{Error: "Invalid Input"})
}
