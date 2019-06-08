/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午2:58
* */
package main

import "github.com/kataras/iris"

func main() {
	app := iris.New()

	html := iris.HTML("./", ".html")
	app.RegisterView(html)
	app.OnErrorCode(iris.StatusNotFound,notFound)

	app.Run(iris.Addr(":8085"),iris.WithCharset("UTF-8"))
}

func notFound(ctx iris.Context)  {
	ctx.View("404.html")
}
