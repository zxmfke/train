package main

import (
	"fmt"
	"github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/context"
)

func main() {

	web.InsertFilter("/bbb/*", web.BeforeRouter, func(ctx *context.Context) {
		fmt.Println("beego middleware")
	})

	bbNS := web.NewNamespace("/bbb")
	bbNS.Get("/b", func(ctx *context.Context) {
		ctx.WriteString("bbb")
	})

	web.AddNamespace(bbNS)

	web.Get("/ping", func(ctx *context.Context) {
		ctx.WriteString("pong")
	})

	web.Run()
}
