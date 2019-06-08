/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午10:11
* */
package controllers

import (
	"github.com/dollarkillerx/iris-xorm-Study/demo2/services"
	"github.com/kataras/iris"
	"github.com/kataras/iris/mvc"
)

type AdminController struct {
	Ctx iris.Context
	Service services.SuperstarService
}

func (c *AdminController) Get() mvc.Result {
	all := c.Service.GetAll()
	return mvc.View{
		Name:"admin/index.html",
		Data:iris.Map{
			"Title":"admin",
			"Datalist":all,
		},
		Layout:"admin/layout.html",
	}
}

func (c *AdminController) GetEdit() mvc.Result {
	id,err := c.Ctx.URLParamInt("id")
	if id != 0 && err == nil{
		get := c.Service.Get(id)
		return mvc.View{
			Name:"admin/edit.html",
			Data:iris.Map{
				"Title":"admin",
				"Info":get,
			},
			Layout:"admin/layout.html",
		}
	}
	return nil
}



