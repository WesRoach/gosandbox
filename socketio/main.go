package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {

	server := socketio.NewServer(nil)

	server.OnConnect("/socket.io/", func(so socketio.Conn) error {
		so.SetContext("")
		fmt.Println("connected:", so.ID())
		return nil
	})

	http.Handle("/socket.io/", server)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
