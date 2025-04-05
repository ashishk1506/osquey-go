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
	resp1.StoreOsVersion()
	fmt.Println(resp1)
	//resp2 := cmd.ExecProgramList()
	//ref2 := &model.ProgramList{}
	//ref2.ClearTable()
	//ref2.StoreProgramList(resp2)
	//fmt.Println(resp2)
	// os version

}
