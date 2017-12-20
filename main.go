package main

import (
	"net/http"
	"log"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello go"))
	})

	http.HandleFunc("/bye", byeFunc)

	log.Println("Starting server... v1")
	log.Fatalln(http.ListenAndServe(":4000", nil))
}

func byeFunc(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("bye go"))
}
