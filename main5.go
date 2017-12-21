package main

import (
	"net/http"
	"log"
	"html/template"
	"time"
	"fmt"
	"strconv"
)

func main() {
	myServer := &http.Server{
		Addr:         ":4000",
		WriteTimeout: 2 * time.Second,
	}
	log.Println("Start server...")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		template, err := template.ParseFiles("main5.tmpl")
		if err != nil {
			log.Fatalf("Parse: %v", err)
		}

		score := r.FormValue("score")
		num, _ := strconv.Atoi(score)

		err = template.Execute(w, num)
		if err != nil {
			fmt.Fprintf(w, "Execute: %v", err)
			return
		}
	})

	log.Fatalln(myServer.ListenAndServe())

}
