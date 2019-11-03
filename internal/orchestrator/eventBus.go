package orchestrator

import (
	"github.com/arielhenryson/dracaena/internal/orchestrator/runnerHandler"
	"github.com/arielhenryson/dracaena/pkg/types"
	"github.com/reactivex/rxgo/observer"
)

func createEventBusObservable(workflow *types.Workflow) observer.Observer {
	return observer.Observer{
		// Register a handler function for every next available item.
		NextHandler: func(data interface{}) {
			msg := data.(types.BusMessage)

			// handle what to do when getting response from any runner
			runnerHandler.RunnerResponseHandler(msg, workflow)
		},

		// Register a handler for any emitted error.
		ErrHandler: func(err error) {
			// fmt.Printf("Encountered error: %v\n", err)
		},

		// Register a handler when a stream is completed.
		DoneHandler: func() {
			workflow.Internal.CompletedJobs++
		},
	}
}
