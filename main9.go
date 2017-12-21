package main

import (
	"gopkg.in/macaron.v1"
	"log"
	"fmt"
)

func main() {
	m := macaron.Classic()
	log.Println("Server starting")
	m.Get("", next1, next2)

	m.Run()
}

func next1(ctx *macaron.Context) string {
	fmt.Println("1 enter")
	ctx.Next()
	fmt.Println("1 exit")
	return ctx.RemoteAddr()
}
func next2(ctx *macaron.Context) string {
	fmt.Println("2 enter")
	ctx.Next()
	fmt.Println("2 exit")
	return ctx.RemoteAddr()
}
