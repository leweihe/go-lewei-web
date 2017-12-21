package main

import (
	"gopkg.in/macaron.v1"
	"log"
	"fmt"
)

func main() {
	m := macaron.Classic()
	log.Println("Server starting")

	m.SetDefaultCookieSecret("default cookie")
	m.Get("/", setCookie, getCookie)
	m.Run()
}

func setCookie(ctx *macaron.Context) string {
	ctx.SetCookie("user", "lewei")
	ctx.SetSecureCookie("pswd", "lewei")
	fmt.Println("set cookie 1")
	ctx.Next()
	fmt.Println("set cookie 2")
	return ""
}
func getCookie(ctx *macaron.Context) string {
	fmt.Println("set cookie 3")
	pswd, _ := ctx.GetSecureCookie("pswd")
	return ctx.GetCookie("user") + " pswd :" + pswd + "."
}
