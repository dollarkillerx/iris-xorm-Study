/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午3:18
* */
package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

func main() {
	app := iris.New()

	// 配置
	mvc.Configure(app.Party("/root"),myMvc)

	app.Run(iris.Addr(":8085"),iris.WithCharset("UTF-8"))
}

func myMvc(app *mvc.Application) {
	app.Handle(new(MyController))
}

// controller
type MyController struct {}

// 再添加路由
func (m *MyController) BeforeActivation(b mvc.BeforeActivation) {
	b.Handle("GET", "/something/{id:long}", "MyCustomHandler",hello)// method,path,funcName,middleware
}

func (m *MyController) Get() string {
	return "Hello World"
}

func (m *MyController) MyCustomHandler(id int64) string {
	return "MyCustomHandler"
}

func hello(ctx iris.Context) {
	fmt.Println("ctx")
	ctx.Next()
}