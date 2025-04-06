package cmd

import (
	"encoding/json"
	"log"
	"os/exec"
	"osquey/model"
)

func ExecOsQueryVersion() model.OsQueryVersion {
	cmd := exec.Command("osqueryi", "--json", "SELECT version FROM osquery_info LIMIT 1;")

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running osqueryi: %v\n", err)
	}

	var verObj []model.OsQueryVersion

	err = json.Unmarshal(output, &verObj)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON output: %v", err)
	}

	if len(verObj) >= 1 {
		return verObj[0]
	}

	return model.OsQueryVersion{}
}
