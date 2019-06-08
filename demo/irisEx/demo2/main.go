/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 上午10:53
* */
package main

import "github.com/kataras/iris"

func main()  {
	app := iris.New()

	// 分组
	userRouter := app.Party("/user")
	// route: /user/{name}/home  例如:/user/dollarKiller/home
	userRouter.Get("/{name:string}/home", func(ctx iris.Context) {
		name := ctx.Params().Get("name")
		ctx.Writef("you name: %s",name)
	})
	// route: /user/post
	userRouter.Post("/post", func(ctx iris.Context) {
		ctx.Writef("method:%s,path;%s",ctx.Method(),ctx.Path())
	})


	app.Run(iris.Addr(":8085"),iris.WithCharset("UTF-8"))
}
