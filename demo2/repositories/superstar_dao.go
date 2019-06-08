/**
* Created by GoLand
* User: dollarkiller
* Date: 19-6-8
* Time: 下午9:14
* */
package repositories

import (
	"errors"
	"github.com/dollarkillerx/iris-xorm-Study/demo2/dbops"
	"github.com/dollarkillerx/iris-xorm-Study/demo2/defs"
)

type SuperstarDao struct {

}

func (s *SuperstarDao) Get(id int) *defs.StarInfo {
	info := &defs.StarInfo{Id: id}
	_, e := dbops.Engine.Get(info)
	if e != nil {
		return nil
	}
	return info
}

func (s *SuperstarDao) GetAll() []*defs.StarInfo {
	infos := make([]*defs.StarInfo,0)
	err := dbops.Engine.Desc("id").Find(infos)
	if err != nil {
		return nil
	}
	return infos
}

func (s *SuperstarDao) Search(country string) []*defs.StarInfo {
	infos := make([]*defs.StarInfo, 0)
	err := dbops.Engine.Where("country=?", country).Desc("id").Find(infos)
	if err != nil {
		return nil
	}
	return infos
}

func (s *SuperstarDao) Delete(id int) error {
	info := &defs.StarInfo{Id: id, SysStatus: 1}
	i, e := dbops.Engine.Id(id).Cols("sys_status").Update(info)
	if e != nil || i ==0 {
		return errors.New("error Update")
	}
	return nil
}

func (s *SuperstarDao) Update(data *defs.StarInfo,columns []string) error {
	i, e := dbops.Engine.Id(data.Id).MustCols(columns...).Update(data)
	if e != nil || i == 0 {
		return errors.New("error Update")
	}
	return nil
}

func (s *SuperstarDao) Create(data defs.StarInfo) error {
	_, e := dbops.Engine.InsertOne(data)
	return e
}


