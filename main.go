package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os/exec"
	"osquey/db"
)

type OsVersion struct {
	Version string `json:"version"`
}

func main() {
	fmt.Println("Hello World")
	db.DB()

	// os version
	cmd := exec.Command("osqueryi", "--json", "SELECT version FROM os_version LIMIT 1;")

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running osqueryi: %v\nOutput: %s", err, output)
	}

	var verObj []OsVersion

	err = json.Unmarshal(output, &verObj)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON output: %v", err)
	}

	fmt.Printf("%+v\n", verObj)

}
