package validator

import (
	"fmt"
	"github.com/arielhenryson/dracaena/pkg/types"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"os"
)

// LoadWorkflow load the main yaml file
func LoadWorkflow(path string) types.Workflow {
	checkWorkFlowPath()

	str := loadFileFromString(path)

	b := []byte(str)
	return convertByteToWorkflowObject(b)
}

// CheckWorkFlowPath check if there is path to workflow file
func checkWorkFlowPath() {
	if len(os.Args) < 2 {
		_, _ = fmt.Fprintf(os.Stderr, "Missing path to workflow file\n")
		os.Exit(1)
	}
}

// loadFileFromString load yaml as string
func loadFileFromString(path string) string {
	b, err := ioutil.ReadFile(path) // just pass the file name
	if err != nil {
		fmt.Print(err)
	}

	str := string(b) // convert content to a 'string'

	return str
}

// convertByteToWorkflowObject convert string to workflow object
func convertByteToWorkflowObject(b []byte) types.Workflow {
	workflow := types.Workflow{}

	err := yaml.Unmarshal(b, &workflow)
	if err != nil {
		log.Fatalf("error: %v", err)
	}

	return workflow
}
