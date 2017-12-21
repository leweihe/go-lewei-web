package main

import (
	"gopkg.in/macaron.v1"
	"log"
	"fmt"
)

func main() {
	macaron.Env = macaron.DEV

	m := macaron.Classic()
	log.Println("Server starting")
	m.Get("",
		func(ctx *macaron.Context) {
			ctx.Data["Num"] = 1

		}, func(ctx *macaron.Context) {
			ctx.Data["Num"] = ctx.Data["Num"].(int) + 1
		}, func(ctx *macaron.Context) {
			ctx.Data["Num"] = ctx.Data["Num"].(int) + 1
		}, func(ctx *macaron.Context) {
			ctx.Data["Num"] = ctx.Data["Num"].(int) + 1
		}, func(ctx *macaron.Context) {
			fmt.Sprintf("Num : %d", ctx.Data["Num"])
		}, func(l log.Logger) {
			l.Fatalln("")
			panic("wrong")
		})

	m.Use(macaron.Static("public"))

	m.Run()
}
