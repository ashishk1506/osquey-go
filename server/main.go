package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"osquey/api"
	"osquey/cmd"
	"osquey/db"
	"osquey/model"
	"time"
)

func main() {
	//fmt.Println("Hello World")
	db.DB()
	go runTask()
	router := gin.Default()
	router.Use(cors.Default())
	router.GET("/latest_data", api.LatestData)
	router.Run(":8081")
}

func runTask() {
	for {
		osQueryVersion := cmd.ExecOsQueryVersion()
		osQueryVersion.ClearTable()
		osQueryVersion.StoreOsQueryVersion()

		osVersion := cmd.ExecOsVersion()
		osVersion.ClearTable()
		osVersion.StoreOsVersion()

		programList := cmd.ExecProgramList()
		program := &model.ProgramList{}
		program.ClearTable()
		program.StoreProgramList(programList)

		time.Sleep(5 * time.Minute)
	}
}
