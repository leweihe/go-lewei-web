package main

import (
	"log"
	"net/http"
	"gopkg.in/macaron.v1"
)

func main() {

	m := macaron.Classic()
	m.Get("/", myHandlerMacaron)
	m.Get("/*", myHandlerMacaron)

	log.Println("Server is running...")
	log.Println(http.ListenAndServe(":4000", m))
}

func myHandlerMacaron(ctx *macaron.Context) string {
	return "the request path is : " + ctx.Req.RequestURI
}