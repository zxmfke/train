package main

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
)

func main() {
	app := iris.New()
	app.Use(iris.Compression)
	app.Use(func(context *context.Context) {
		context.Next()
		return
	})

	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("hello world!")
	})

	app.Listen(":6666")
}
