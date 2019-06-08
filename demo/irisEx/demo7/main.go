/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午4:22
* */
package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	recover2 "github.com/kataras/iris/middleware/recover"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()

	app.Use(recover2.New()) // 恐慌恢复
	app.Use(logger.New()) // 输入到终端

	mvc.New(app).Handle(new(Container))

	app.Run(iris.Addr(":8085"),iris.WithCharset("UTF-8"))
}

type Container struct {}

func (c *Container) Get() mvc.Result {
	return mvc.Response{
		ContentType:"text/html",
		Text:"<h1>Welcome</h1>",
	}
}

func (c *Container) GetPing() string {
	return "ping"
}


func (c *Container) GetHello() interface{} {
	return map[string]string{
		"message":"Hello World",
	}
}

func (c *Container) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET","/hello/{name:string}","Hello")
}

func (c *Container) Hello(name string) string {
	return "Hello World " + name
}

