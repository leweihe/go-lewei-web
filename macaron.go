package main

import (
	"github.com/Unknwon/macaron"
	"log"
	"net/http"
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