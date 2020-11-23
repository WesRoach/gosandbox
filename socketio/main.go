package main

import (
	"fmt"
	"log"
	"net/http"

	socketio "github.com/googollee/go-socket.io"
)

func main() {

	server := socketio.NewServer(nil)

	server.OnConnect("/socket.io/", func(s socketio.Conn) error {
		s.SetContext("")
		fmt.Println("connected:", s.ID())
		return nil
	})

	http.Handle("/socket.io/", server)
	log.Fatal(http.ListenAndServe(":5000", nil))
}
