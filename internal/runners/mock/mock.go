package mock

import (
	"github.com/arielhenryson/dracaena/pkg/types"
	"github.com/arielhenryson/dracaena/pkg/types/runnerResponseTypes"
	"github.com/reactivex/rxgo/observer"
	"time"
)

// Run will start a job on Mock runner
func Run(workflowID string, id string, job *types.WorkflowJobRunner, observable observer.Observer) {
	time.Sleep(1 * time.Second)

	observable.NextHandler(
		types.BusMessage{
			JobId:   id,
			Type:    runnerResponseTypes.Log,
			Payload: job.Payload.Commands[0],
		})

	observable.NextHandler(
		types.BusMessage{
			JobId:   id,
			Type:    runnerResponseTypes.Exit,
			Payload: "1",
		})

	observable.DoneHandler()
}
