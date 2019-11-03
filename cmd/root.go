package cmd

import (
	"github.com/arielhenryson/dracaena/internal/orchestrator"
	"github.com/arielhenryson/dracaena/pkg/logger"
	"github.com/arielhenryson/dracaena/pkg/validator"
	"os"
)

// Execute run the main app flow
func Execute() {
	workflow := validator.LoadWorkflow(os.Args[1])

	logger.Log("Starting " + workflow.Name)

	orchestrator.Execute(workflow)
}