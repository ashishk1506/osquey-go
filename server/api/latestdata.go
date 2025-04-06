package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"osquey/model"
)

func LatestData(c *gin.Context) {

	osQueryData := &model.OsQueryVersion{}
	osQueryData.LoadOsQueryVersion()
	fmt.Println("read osQueryVersion", osQueryData)

	osData := &model.OsVersion{}
	osData.LoadOsVersion()
	fmt.Println("read osVersion", osData)

	var programListReadData []model.ProgramList
	new(model.ProgramList).LoadProgramList(&programListReadData)
	fmt.Println("read data list is", programListReadData)

	c.IndentedJSON(http.StatusOK, gin.H{
		"os_version":      osQueryData.Version,
		"osquery_version": osData.Version,
		"program_list":    programListReadData,
	})
}
