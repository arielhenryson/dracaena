package registerJobRunners

import "github.com/arielhenryson/dracaena/pkg/types"

// calculateDependenciesJobs set number of job to complete before start job
func calculateDependenciesJobs(jobId string ,workflow *types.Workflow) int  {
	jobToCompleteBeforeStart := 0

	for id := range workflow.Jobs {
		if id == jobId {
			return jobToCompleteBeforeStart
		}

		if !workflow.Jobs[id].Async {
			jobToCompleteBeforeStart++
		}
	}

	return jobToCompleteBeforeStart
}