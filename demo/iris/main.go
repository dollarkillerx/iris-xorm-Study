/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 上午9:48
* */
package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	htmlEngine := iris.HTML("./view", ".html")
	app.RegisterView(htmlEngine)


	app.Get("/", func(ctx iris.Context) {
		ctx.WriteString("Hello World Iris")
	})

	app.Get("hello", func(ctx iris.Context) {
		ctx.ViewData("Title","测试页面")
		ctx.ViewData("Content","Hello World Iris")
		ctx.View("Hello.html")
	})

	app.Run(iris.Addr(":8085"),iris.WithCharset("UTF-8"))
}
