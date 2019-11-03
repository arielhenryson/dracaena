package runnerHandler

import (
	"github.com/arielhenryson/dracaena/pkg/types"
	"github.com/arielhenryson/dracaena/pkg/types/jobStatusTypes"
)

func runnerExitHandler(msg types.BusMessage, workflow *types.Workflow)  {
	workflow.Jobs[msg.JobId].Internal.Status = jobStatusTypes.Done
}