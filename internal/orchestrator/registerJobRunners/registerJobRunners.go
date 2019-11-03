package registerJobRunners

import (
	"github.com/arielhenryson/dracaena/pkg/types"
	"github.com/reactivex/rxgo/observer"
)

// register job handler for every job
func RegisterJobRunners(workflow *types.Workflow, eventBusObservable observer.Observer)  {
	for id := range workflow.Jobs {
		workflow.Jobs[id].Internal.JobsToCompleteBeforeStart = calculateDependenciesJobs(id, workflow)

		runJob(id, workflow, eventBusObservable)
	}
}
