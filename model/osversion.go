package model

type OsVersion struct {
	Id      int32  `json:"id" orm:"column(id);pk;auto"`
	Version string `json:"version" orm:"column(version)"`
}
