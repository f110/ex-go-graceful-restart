package main

import (
	"fmt"
	"github.com/lestrrat/go-server-starter/listener"
	"log"
	"net"
	"net/http"
	"os"
)

func hello(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintf(w, "Hello world!")
}

func newHandler() http.Handler {
	h := http.NewServeMux()
	h.HandleFunc("/", hello)
	return h
}

func main()  {
	var l net.Listener

	if os.Getenv("SERVER_STARTER_PORT") != "" {
		listeners, err := listener.ListenAll()
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			return
		}

		if len(listeners) > 0 {
			l = listeners[0]
		}
	}

	log.Printf("Start Go HTTP Server")
	log.Fatal(http.Serve(l, newHandler()))
}
