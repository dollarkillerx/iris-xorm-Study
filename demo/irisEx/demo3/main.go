/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午12:07
* */
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"strconv"
)

func main() {
	app := iris.New()

	// 路由传参
	app.Get("/username/{name}", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		fmt.Println(name)
	})

	// 设置参数
	app.Get("/profile/{id:int min(1)}", func(ctx iris.Context) {
		i, e := ctx.Params().GetInt("id")
		if e != nil {
			ctx.WriteString("error you input")
		}

		ctx.WriteString(strconv.Itoa(i))
	})

	// 设置错误码
	app.Get("/profile/{id:int min(1)}/friends/{friendid:int max(8) else 504}", func(ctx iris.Context) {
		i, _ := ctx.Params().GetInt("id")
		getInt, _ := ctx.Params().GetInt("friendid")
		ctx.Writef("Hello id:%d looking for friend id: ",i,getInt)
	})// 如果没有传递所有路由的macros，这将抛出504错误代码而不是404.

	// 正则表达式
	app.Get("/lowercase/{name:string regexp(^[a-z]+)}", func(ctx iris.Context) {
		ctx.Writef("name should be only lowercase, otherwise this handler will never executed: %s", ctx.Params().Get("name"))
	})

	app.Run(iris.Addr(":8085"),iris.WithCharset("UTF-8"))
}
