package runCommand

import (
	"bytes"
	"context"
	"io/ioutil"

	"github.com/arielhenryson/dracaena/internal/runners/docker/helpers/createCommand"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
	"github.com/docker/docker/pkg/stdcopy"
)

// RunCommand run command on container
func RunCommand(ctx context.Context, cli *client.Client, containerID string, command string) (ExecResult, error) {
	exeID, _ := prepareCommand(ctx, cli, containerID, command)

	return inspectExecResp(ctx, cli, exeID)
}

func prepareCommand(ctx context.Context, cli *client.Client, containerID string, cmd string) (string, error) {
	command := createCommand.CreateCommand(cmd)

	config := types.ExecConfig{
		AttachStderr: true,
		AttachStdout: true,
		Cmd:          command,
		Detach:       false,
	}

	commandResponse, err := cli.ContainerExecCreate(ctx, containerID, config)
	if err != nil {
		panic(err)
	}

	return commandResponse.ID, err
}

func inspectExecResp(ctx context.Context, cli *client.Client, id string) (ExecResult, error) {
	var execResult ExecResult

	aRes, err := cli.ContainerExecAttach(ctx, id, types.ExecConfig{})
	if err != nil {
		return execResult, err
	}
	defer aRes.Close()

	// read the output
	var outBuf, errBuf bytes.Buffer
	outputDone := make(chan error)

	go func() {
		// StdCopy demultiplexes the stream into two buffers
		_, err = stdcopy.StdCopy(&outBuf, &errBuf, aRes.Reader)
		outputDone <- err
	}()

	select {
	case err := <-outputDone:
		if err != nil {
			return ExecResult{}, err
		}
		break

	case <-ctx.Done():
		return ExecResult{}, ctx.Err()
	}

	stdout, err := ioutil.ReadAll(&outBuf)
	if err != nil {
		return execResult, err
	}
	stderr, err := ioutil.ReadAll(&errBuf)
	if err != nil {
		return execResult, err
	}

	// get the exit code
	iRes, err := cli.ContainerExecInspect(ctx, id)
	if err != nil {
		return ExecResult{}, err
	}

	return ExecResult{
		ExitCode: iRes.ExitCode,
		StdOut:   string(stdout),
		StdErr:   string(stderr),
	}, nil
}

// ExecResult represents a result returned from Exec()
type ExecResult struct {
	StdOut   string
	StdErr   string
	ExitCode int
}
