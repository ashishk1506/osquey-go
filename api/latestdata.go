package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"osquey/cmd"
	"osquey/model"
)

func LatestData(c *gin.Context) {

	osQueryVersion := cmd.ExecOsQueryVersion()
	osQueryVersion.ClearTable()
	osQueryVersion.StoreOsQueryVersion()
	osQueryData := &model.OsQueryVersion{}
	osQueryData.LoadOsQueryVersion()
	fmt.Println("read osQueryVersion", osQueryData)

	osVersion := cmd.ExecOsVersion()
	osVersion.ClearTable()
	osVersion.StoreOsVersion()
	osData := &model.OsVersion{}
	osData.LoadOsVersion()
	fmt.Println("read osVersion", osData)

	programList := cmd.ExecProgramList()
	program := &model.ProgramList{}
	program.ClearTable()
	program.StoreProgramList(programList)
	var programListReadData []model.ProgramList
	new(model.ProgramList).LoadProgramList(&programListReadData)
	fmt.Println("read data list is", programListReadData)

	c.IndentedJSON(http.StatusOK, gin.H{
		"os_version":      osQueryData.Version,
		"osquery_version": osData.Version,
		"program_list":    programListReadData,
	})
}
