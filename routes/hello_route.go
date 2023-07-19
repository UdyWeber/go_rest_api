package routes

import (
	"fmt"
	"net/http"
	"strings"
	"udyweber/go_rest_api/router"
)

func Hello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	authHeader := r.Header.Get("Authorization")

	if authHeader == "" {
		http.Error(w, "Malformed Token", http.StatusBadRequest)
		return
	}

	if !strings.HasPrefix(authHeader, "Bearer ") {
		http.Error(w, "Token Must be a Bearer", http.StatusBadRequest)
		return
	}

	token := strings.TrimPrefix(authHeader, "Bearer ")
	user, ok := router.GlobalState.Get(token)
	
	if !ok {
		http.Error(w, "User Not Found", http.StatusBadRequest)
		return
	}
	
	fmt.Fprintf(w, "Welcome friend: %s", user.Name)
}