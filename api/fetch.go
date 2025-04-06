package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"osquey/cmd"
	"osquey/model"
)

func Fetch(c *gin.Context) {

	osQueryVersion := cmd.ExecOsQueryVersion()
	osQueryVersion.ClearTable()
	osQueryVersion.StoreOsQueryVersion()
	osQueryVersion.LoadOsQueryVersion()
	fmt.Println("read osQueryVersion", osQueryVersion)

	osVersion := cmd.ExecOsVersion()
	osVersion.ClearTable()
	osVersion.StoreOsVersion()
	osVersion.LoadOsVersion()
	fmt.Println("read osVersion", osVersion)

	programList := cmd.ExecProgramList()
	program := &model.ProgramList{}
	program.ClearTable()
	program.StoreProgramList(programList)
	//var programListReadData []model.ProgramList
	//program.LoadProgramList(programListReadData)
	fmt.Println("read data list is", programList)

}
