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
	//resp1 := cmd.ExecOsVersion()
	//fmt.Println(resp1)
	resp2 := cmd.ExecProgramList()
	fmt.Println(resp2)
	// os version

}
