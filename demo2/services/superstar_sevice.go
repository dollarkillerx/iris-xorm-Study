/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午9:54
* */
package services

import "github.com/dollarkillerx/iris-xorm-Study/demo2/defs"

type SuperstarService interface {
	GetAll() []*defs.StarInfo
	Get(id int) *defs.StarInfo
	Delete(id int) error
	Update(data *defs.StarInfo,columns []string) error
	Create(data defs.StarInfo) error
	Search(country string) []*defs.StarInfo
}
