package main

import (
	"gopkg.in/macaron.v1"
	"log"
)

func main() {
	m := macaron.Classic()
	log.Println("Server starting")
	m.Get("",
		func(ctx *macaron.Context) string {
			return ctx.Query("uid")
		})

	m.Run()
}
