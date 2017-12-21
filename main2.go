package main

import (
	"net/http"
	"log"
	"os"
	"os/signal"
	"time"
)

func main() {
	myServer := &http.Server{
		Addr:         ":4000",
		WriteTimeout: 2 * time.Second,
	}
	mux := http.NewServeMux()
	mux.Handle("/", &myHandler{})
	mux.HandleFunc("/bye", sayBye)

	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)

	go func() {
		<-quit
		if err := myServer.Close(); err != nil {
			log.Fatal("Close server: ", err)
		}
	}()

	log.Println("Starting server... v2")
	err := myServer.ListenAndServe()
	if err != nil {
		if err == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected", err)
		}
	}
	log.Println("Server exit")
}

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi go... v2. URL" + r.URL.String()))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	time.Sleep(3 * time.Second)
	w.Write([]byte("bye go... v2"))
}
