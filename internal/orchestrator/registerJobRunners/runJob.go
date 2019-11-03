package registerJobRunners

import (
	"fmt"
	"github.com/arielhenryson/dracaena/internal/orchestrator/config"
	"github.com/arielhenryson/dracaena/internal/runners/docker"
	"github.com/arielhenryson/dracaena/internal/runners/mock"
	"github.com/arielhenryson/dracaena/pkg/types"
	"github.com/arielhenryson/dracaena/pkg/types/jobRunnerTypes"
	"github.com/reactivex/rxgo/observer"
	"time"
)

// runJob will run a new job
func runJob(id string, workflow *types.Workflow, eventBusObservable observer.Observer)  {
	config := config.MainConfig()
	job := workflow.Jobs[id]

	workflowID := workflow.Internal.ID

	// check if we need to wait before starting this job
	wait := job.Internal.JobsToCompleteBeforeStart > workflow.Internal.CompletedJobs
	if wait {
		fmt.Println("wait for job to complete ")
		time.Sleep(config.JobWaitInterval * time.Millisecond)
		runJob(id, workflow, eventBusObservable)

		return
	}

	// select which runner will run the job
	switch job.Runner {
	case jobRunnerTypes.Mock:
		go mock.Run(workflowID, id, job, eventBusObservable)
		break
	case jobRunnerTypes.Docker:
		go docker.Run(workflowID, id, job, eventBusObservable, workflow.Name)
		break
	}
}