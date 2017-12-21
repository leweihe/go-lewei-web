package main

import (
	"gopkg.in/macaron.v1"
	"log"
	"fmt"
)

func main() {
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
		}, func(ctx *macaron.Context) string {
			return fmt.Sprintf("Num : %d", ctx.Data["Num"])
		})

	m.Run()
}
