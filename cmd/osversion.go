package cmd

import (
	"encoding/json"
	"log"
	"os/exec"
)

type OsVersion struct {
	Version string `json:"version"`
}

func ExecOsVersion() string {
	cmd := exec.Command("osqueryi", "--json", "SELECT version FROM os_version LIMIT 1;")

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running osqueryi: %v\n", err)
	}

	var verObj []OsVersion

	err = json.Unmarshal(output, &verObj)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON output: %v", err)
	}

	if len(verObj) >= 1 {
		return verObj[0].Version
	}

	return ""
}
