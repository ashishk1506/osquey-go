package main

import (
	"fmt"
	"osquey/cmd"
	"osquey/db"
)

func main() {
	fmt.Println("Hello World")
	db.DB()
	//resp := cmd.ExecOsQueryVersion()
	//fmt.Println(resp)
	resp1 := cmd.ExecOsVersion()
	fmt.Println(resp1)
	// os version

}
