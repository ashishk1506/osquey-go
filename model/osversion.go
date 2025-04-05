package model

import (
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
)

type OsVersion struct {
	Id      int32  `json:"id" orm:"column(id);pk;auto"`
	Version string `json:"version" orm:"column(version)"`
}

func (p *OsVersion) TableName() string {
	return "os_version"
}

func (p *OsVersion) ClearTable() {
	orm := orm.NewOrm()
	count, err := orm.Raw("TRUNCATE TABLE ?", p.TableName()).Exec()
	if err != nil {
		log.Error("Error clearing table:%v", err)
	}
	log.Infof("Number of records deleted: %d\n", count)
}

func (p *OsVersion) StoreOsVersion() {
	orm := orm.NewOrm()
	if p.Version == "" {
		return
	}
	count, err := orm.Insert(p)
	if err != nil {
		log.Error("Error running store program list insert: %v", err)
	}
	log.Infof("Number of records inserted: %d\n", count)

}
