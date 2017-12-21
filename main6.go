package main

import (
	"net/http"
	"log"
	"html/template"
	"time"
	"fmt"
)

func main() {
	myServer := &http.Server{
		Addr:         ":4000",
		WriteTimeout: 2 * time.Second,
	}
	log.Println("Start server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("main6.tmpl")
		if err != nil {
			log.Fatalf("Parse: %v", err)
		}

		err = template.Execute(w, map[string]interface{}{
			"Request": r,
			"Score":   99,
		})
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Fatalln(myServer.ListenAndServe())

}
