package orchestrator

import (
	"github.com/arielhenryson/dracaena/internal/orchestrator/config"
	"github.com/arielhenryson/dracaena/pkg/types"
	"github.com/arielhenryson/dracaena/pkg/types/jobStatusTypes"
	"time"
)

// cycle throw the jobs to determine what is the status of the workflow
func lifecycleManager(workflow *types.Workflow)  {
	config := config.MainConfig()

	workIsDone := false
	for !workIsDone {
		workIsDone = true

		for k := range workflow.Jobs  {
			job := workflow.Jobs[k]

			if job.Internal.Status != jobStatusTypes.Done {
				workIsDone = false
			}
		}

		time.Sleep(config.LifeCycleInterval * time.Millisecond)
	}
}
