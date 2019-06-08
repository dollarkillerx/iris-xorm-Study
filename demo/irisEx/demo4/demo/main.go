/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午12:41
* */
package main

import (
	"fmt"
	"github.com/kataras/iris"
)

func main() {
	app := iris.New()

	// 注册前置全局中间件
	app.Use(before)
	// 主持后置
	app.Done(after)

	app.Get("/", func(ctx iris.Context) {
		ctx.HTML("<h1>Hello</h1>")
		ctx.Next()
	})

	app.Run(iris.Addr(":8085"),iris.WithCharset("UTF-8"))
}

func before(ctx iris.Context)  {
	header := ctx.GetHeader("token")
	fmt.Println("全局前置..........",header)
	ctx.Next()
}

func after(ctx iris.Context)  {
	fmt.Println("后置............")
}