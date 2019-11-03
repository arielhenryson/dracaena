package runnerHandler

import (
	"github.com/arielhenryson/dracaena/pkg/logger"
	"github.com/arielhenryson/dracaena/pkg/types"
	"github.com/arielhenryson/dracaena/pkg/types/runnerResponseTypes"
)

// runnerResponseHandler handle what to do when getting response from any runner
func RunnerResponseHandler(msg types.BusMessage, workflow *types.Workflow)  {
	if msg.Type == runnerResponseTypes.Exit {
		runnerExitHandler(msg, workflow)

		return
	}


	if msg.Type == runnerResponseTypes.Log {
		logger.Log(msg.Payload)
	}
}