package middleware

import (
	"fmt"
	"net/http"
)

// Home function
func Home(w http.ResponseWriter, r *http.Request) {
	// Get method, Welcome screen
	fmt.Fprintln(w, "Welcome to our bank")
}
