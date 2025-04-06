package model

import (
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
)

type OsQueryVersion struct {
	Id      int32  `json:"id" orm:"column(id);pk;auto"`
	Version string `json:"version" orm:"column(version)"`
}

func (p *OsQueryVersion) TableName() string {
	return "os_query_version"
}

func (p *OsQueryVersion) ClearTable() {
	orm := orm.NewOrm()
	count, err := orm.Raw("TRUNCATE TABLE " + p.TableName()).Exec()
	if err != nil {
		log.Errorf("Error clearing table:%v", err)
	}
	log.Infof("Number of records deleted: %d\n", count)
}

func (p *OsQueryVersion) StoreOsQueryVersion() {
	orm := orm.NewOrm()
	if p.Version == "" {
		return
	}
	count, err := orm.Insert(p)
	if err != nil {
		log.Error("Error running store osquery insert: %v", err)
	}
	log.Infof("Number of records inserted: %d\n", count)

}

func (p *OsQueryVersion) LoadOsQueryVersion() {
	orm := orm.NewOrm()
	//OsQueryObj := new(OsQueryVersion)
	err := orm.Read(p)
	if err != nil {
		log.Error("Error reading osquery insert: %v", err)
	}
}
