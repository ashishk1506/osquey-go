package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"osquey/api"
	"osquey/db"
)

func main() {
	fmt.Println("Hello World")
	db.DB()
	router := gin.Default()
	router.GET("/fetch", api.Fetch)
	router.Run(":8081")
}
