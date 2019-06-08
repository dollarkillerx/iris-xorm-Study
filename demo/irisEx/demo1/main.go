/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 上午10:45
* */
package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()
	//GET 方法
	app.Get("/", handler)
	// POST 方法
	app.Post("/", handler)
	// PUT 方法
	app.Put("/", handler)
	// DELETE 方法
	app.Delete("/", handler)
	//OPTIONS 方法
	app.Options("/", handler)
	//TRACE 方法
	app.Trace("/", handler)
	//CONNECT 方法
	app.Connect("/", handler)
	//HEAD 方法
	app.Head("/", handler)
	// PATCH 方法
	app.Patch("/", handler)
	//任意的http请求方法如option等
	app.Any("/", handler)

	app.Run(iris.Addr(":8085"),iris.WithCharset("UTF-8"))
}

// 处理函数
func handler(ctx iris.Context) {
	ctx.Writef("methdo:%s path:%s",ctx.Method(),ctx.Path())
}