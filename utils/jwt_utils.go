package utils

import (
	"net/http"
	"strings"
)

func ExtractTokenFromHeader(req *http.Request) string {
	// Get the Authorization header
	bearerToken := req.Header.Get("Authorization")

	// Check if the header is empty or doesn't contain "Bearer "
	if bearerToken == "" || !strings.HasPrefix(bearerToken, "Bearer ") {
		return ""
	}

	// Extract the token from the "Bearer " prefix
	return strings.TrimPrefix(bearerToken, "Bearer ")
}
