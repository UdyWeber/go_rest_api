package routes

import (
	"fmt"
	"net/http"
)

func RootRoute(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!\nRoot Route here :D")
}
