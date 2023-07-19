package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"udyweber/go_rest_api/routes"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", routes.RootRoute)
	mux.HandleFunc("/user/create", routes.CreateUser)
	mux.HandleFunc("/user/all", routes.GetAllUsers)
	mux.HandleFunc("/hello", routes.Hello)

	fmt.Println("Ligou na porta :8213")

	server := &http.Server{
		Addr: ":8213",
		Handler: mux,
		BaseContext: func(l net.Listener) context.Context {
			ctx := context.Background()
			ctx = context.WithValue(ctx, "serverAddr", l.Addr().String())
			return ctx
		},
	}

	http_err := server.ListenAndServe()
	log.Fatal(http_err)
}
