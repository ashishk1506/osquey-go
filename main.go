package main

import (
	"fmt"
	"osquey/cmd"
	"osquey/db"
	"osquey/model"
)

func main() {
	fmt.Println("Hello World")
	db.DB()
	//resp := cmd.ExecOsQueryVersion()
	//fmt.Println(resp)
	//resp1 := cmd.ExecOsVersion()
	//fmt.Println(resp1)
	resp2 := cmd.ExecProgramList()
	ref := &model.ProgramList{}
	ref.ClearTable()
	ref.GetProgramList(resp2)
	fmt.Println(resp2)
	// os version

}
