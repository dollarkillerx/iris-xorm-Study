/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午12:26
* */
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"strings"
)

func main() {
	app := iris.New()

	app.Get("/name/{name}",before,mainHandler,after)
	// before,mainHandler,after 感觉这个设计思路好像那个koa2
	// 前一个通过ctx.Next() 进入下一个方法

	app.Run(iris.Addr(":8085"),iris.WithCharset("UTF-8"))
}

func before(ctx iris.Context)  {
	name := ctx.Params().Get("name")
	if strings.EqualFold(name,"dollarkiller") {
		fmt.Println("before...............")
		ctx.Next()
		return
	}
	ctx.WriteString("error none")
}

func after(ctx iris.Context) {
	fmt.Println("after.....................")
}

func mainHandler(ctx iris.Context) {
	fmt.Println("main.................")
	ctx.WriteString("ok........")
	ctx.Next()
}