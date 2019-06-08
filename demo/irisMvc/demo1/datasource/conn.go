/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午5:13
* */
package datasource

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

var (
	Engine *xorm.Engine
	e error
)

// 初始化数据源
func init() {
	Engine, e = xorm.NewEngine("mysql", "dsn")
	if e != nil {
		panic(e.Error())
	}
}
