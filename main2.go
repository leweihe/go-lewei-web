package main

import (
	"net/http"
	"log"
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

	log.Println("Starting server... v2")
	log.Fatal(myServer.ListenAndServe())
}

type myHandler struct {
}

func (*myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hi go... v2. URL" + r.URL.String()))
}

func sayBye(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye go... v2"))
}
