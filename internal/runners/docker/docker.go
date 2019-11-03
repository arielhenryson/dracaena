package docker

import (
	"context"
	"github.com/arielhenryson/dracaena/internal/runners/docker/api/createClient"
	"github.com/arielhenryson/dracaena/internal/runners/docker/api/createContainer"
	"github.com/arielhenryson/dracaena/internal/runners/docker/api/createNetwork"
	"github.com/arielhenryson/dracaena/internal/runners/docker/api/createVolume"
	"github.com/arielhenryson/dracaena/internal/runners/docker/api/pullImage"
	"github.com/arielhenryson/dracaena/pkg/types"
	"github.com/arielhenryson/dracaena/pkg/types/runnerResponseTypes"
	"github.com/reactivex/rxgo/observer"
)

// Run will start a job on Docker runner
func Run(workflowID string, id string, job *types.WorkflowJobRunner, observable observer.Observer, workflowName string) {
	ctx := context.Background()
	cli, _ := createClient.CreateClient()


	volName, _ := createVolume.CreateVolume(cli, ctx, workflowName)

	networkID, _ := createNetwork.CreateNetwork(cli, ctx, workflowName)

	pullImage.PullImage(ctx, cli, job.Payload.Image)

	_, _ = createContainer.CreateContainer(cli, ctx, job.Payload.Image, volName,job.Payload.Commands, id, networkID)

	observable.NextHandler(
		types.BusMessage{
			JobId:   id,
			Type:    runnerResponseTypes.Exit,
			Payload: "0",
		})

	observable.DoneHandler()
}
