/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午10:02
* */
package controllers

import (
	"github.com/dollarkillerx/iris-xorm-Study/demo2/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type IndexController struct {
	Ctx iris.Context
	Service services.SuperstarService
}

func (c *IndexController) Get() mvc.Result {
	all := c.Service.GetAll()
	return mvc.View{
		Name:"index.html",
		Data:iris.Map{
			"Title":"Home",
			"Datalist":all,
		},
	}
}

func (c *IndexController) GetBy(id int) mvc.Result {
	if id <1 {
		return mvc.Response{
			Path:"/",
		}
	}
	get := c.Service.Get(id)
	return mvc.View{
		Name:"info.html",
		Data:iris.Map{
			"Title":"Info",
			"Data":get,
		},
	}
}

func (c *IndexController) GetSearch() mvc.Result {
	country := c.Ctx.URLParam("country")
	search := c.Service.Search(country)
	return mvc.View{
		Name:"index.html",
		Data:iris.Map{
			"Title":"search",
			"Datalist":search,
		},
	}
}