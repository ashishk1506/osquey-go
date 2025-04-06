package cmd

import (
	"encoding/json"
	"log"
	"os/exec"
	"osquey/model"
)

func ExecProgramList() []model.ProgramList {
	cmd := exec.Command("osqueryi", "--json", "select name, version, install_date, install_location from programs")

	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("Error running osqueryi: %v\n", err)
	}

	var programList []model.ProgramList

	err = json.Unmarshal(output, &programList)
	if err != nil {
		log.Fatalf("Error unmarshalling JSON output: %v", err)
	}

	return programList

}
