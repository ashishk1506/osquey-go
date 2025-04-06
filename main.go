package main

import (
	"github.com/gin-gonic/gin"
	"osquey/api"
	"osquey/db"
)

func main() {
	//fmt.Println("Hello World")
	db.DB()
	router := gin.Default()
	router.GET("/latest_data", api.LatestData)
	router.GET("/home", api.Home)
	router.Run(":8081")
}
