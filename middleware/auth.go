package middleware

import (
	"net/http"
)

// AdminOnly is a "Wrapper" function.
// It takes a Handler, adds security logic, and returns a secure Handler.
func AdminOnly(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// 1. Get the username/password from the request header
		user, pass, ok := r.BasicAuth()

		// 2. Validate Credentials (Hardcoded for now)
		// In a real app, you would check this against your database!
		if !ok || user != "admin" || pass != "secret123" {
			// If wrong, tell the browser to ask for login
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// 3. If correct, run the original handler
		next(w, r)
	}
}
