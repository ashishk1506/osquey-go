package model

import (
	"github.com/astaxie/beego/orm"
	"github.com/prometheus/common/log"
)

type ProgramList struct {
	Id              int32  `json:"id" orm:"column(id);pk;auto"`
	Name            string `json:"name" orm:"column(name)`
	Version         string `json:"version" orm:"column(version)"`
	InstallDate     string `json:"install_date" orm:"column(install_date)"`
	InstallLocation string `json:"install_location" orm:"column(install_location)"`
}

func (p *ProgramList) ClearTable() {
	orm := orm.NewOrm()
	count, err := orm.Raw("TRUNCATE TABLE program_list").Exec()
	if err != nil {
		log.Error("Error clearing table:", err)
	}
	log.Infof("Number of records deleted: %d\n", count)
}

func (p *ProgramList) StoreProgramList(data []ProgramList) {
	orm := orm.NewOrm()
	count, err := orm.InsertMulti(len(data), data)
	if err != nil {
		log.Error("Error running store program list insert: %v", err)
	}
	log.Infof("Number of records inserted: %d\n", count)

}
