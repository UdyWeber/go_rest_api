package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"net/http"
	"udyweber/go_rest_api/routes"

	"github.com/gorilla/websocket"
	"github.com/rs/cors"
)

var upgrader = websocket.Upgrader {
	ReadBufferSize: 1024,
	WriteBufferSize: 1024,
}

func reader(conn *websocket.Conn){
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Printf("Error reading message: %s", err)
			return
		}

		fmt.Printf("Message retrieved: %s\n", p)

		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Printf("Error writing message: %s", err)
			return
		}
	}
}


func wsEndpoint(w http.ResponseWriter, r *http.Request) {
    upgrader.CheckOrigin = func(r *http.Request) bool { return true}

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Printf("Error upgrading websocket %s\n", err)
	}

	log.Println("Client connected!")
	err = ws.WriteMessage(1, []byte("Hi Client!"))

	if err != nil {
		log.Println(err)
	}

	reader(ws)
}

func setupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", routes.RootRoute)
	mux.HandleFunc("/ws", wsEndpoint)
	mux.HandleFunc("/user/create", routes.CreateUser)
	mux.HandleFunc("/user/all", routes.GetAllUsers)
	mux.HandleFunc("/hello", routes.Hello)
}

func main() {
	mux := http.NewServeMux()
	handler := cors.Default().Handler(mux)

	setupRoutes(mux)

	fmt.Println("Ligou na porta :8213")

	server := &http.Server{
		Addr: ":8213",
		Handler: handler,
		BaseContext: func(l net.Listener) context.Context {
			ctx := context.Background()
			ctx = context.WithValue(ctx, "serverAddr", l.Addr().String())
			return ctx
		},
	}

	http_err := server.ListenAndServe()
	log.Fatal(http_err)
}
