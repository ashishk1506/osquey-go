package model

type ProgramList struct {
	Name            string `json:"name";orm:"column(name);size(255)"`
	Version         string `json:"version";orm:"column(version);size(255)"`
	InstallDate     string `json:"install_date";orm:"column(install_date);size(255)"`
	InstallLocation string `json:"install_location";orm:"column(install_location);size(1024)"`
}

type OsVersion struct {
	Version string `json:"version"`
}

type OsQueryVersion struct {
	Version string `json:"version"`
}
