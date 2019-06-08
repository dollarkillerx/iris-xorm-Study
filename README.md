# iris-xorm-Study
iris xorm Study

文档地址:https://www.kancloud.cn/adapa/go-web-iris
### init
``` 
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
```

### 模板视图
``` 
	htmlEngine := iris.HTML("./view", ".html") // 模板引擎 第一个是模板地址
	app.RegisterView(htmlEngine) // 注册
	
	app.Get("hello", func(ctx iris.Context) {
    		ctx.ViewData("Title","测试页面")
    		ctx.ViewData("Content","Hello World Iris")
    		ctx.View("Hello.html")
    })
```

### MVC
目录
```
.
├── config  配置
├── datamodels  定义
│   └── movie.go
├── datasource 数据源
│   └── conn.go 
├── repositories dao操作
├── services 服务层
└── web 
    ├── controllers  控制器
    ├── main.go
    ├── middleware 
    └── views
        └── hello
```