/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午6:07
* */
package dbops

import (
	"github.com/dollarkillerx/iris-xorm-Study/demo2/config"
	"github.com/go-xorm/xorm"
)

var (
	Engine *xorm.Engine
	e error
)

func init() {
	Engine,e = xorm.NewEngine(config.Configs.DriverName,config.Configs.Dsn)
	if e != nil {
		panic(e.Error())
	}
	if config.Configs.Debug {
		Engine.ShowSQL(true)
	}
}
