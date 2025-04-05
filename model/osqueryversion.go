package model

type OsQueryVersion struct {
	Id      int32  `json:"id" orm:"column(id);pk;auto"`
	Version string `json:"version" orm:"column(version)"`
}
