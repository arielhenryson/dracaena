package orchestrator

import (
	"github.com/arielhenryson/dracaena/internal/orchestrator/registerJobRunners"
	"github.com/arielhenryson/dracaena/pkg/types"
	"github.com/google/uuid"
)

// Execute will run the workflow
func Execute(workflow types.Workflow)  {
	eventBusObservable := createEventBusObservable(&workflow)

	workflow.Internal.ID = uuid.New().String()


	workflow.Internal.CompletedJobs = 0

	// register runner for every job
	registerJobRunners.RegisterJobRunners(&workflow, eventBusObservable)

	// cycle throw the jobs to determine what is the status of the workflow
	lifecycleManager(&workflow)
}


