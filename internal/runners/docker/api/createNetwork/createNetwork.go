package createNetwork

import (
	"context"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

// CreateNetwork create network so container in this network will be able to talk
func CreateNetwork(cli *client.Client, ctx context.Context, networkName string) (string, error)  {
	// if we already have a network with this name we don't want to crate a new one
	networkExist, networkID := isNetworkExists(cli, ctx, networkName)
	if networkExist {
		return networkID, nil
	}

	res, err := cli.NetworkCreate(ctx, networkName, types.NetworkCreate{})

	return res.ID, err
}

// check if network already created with this networkName
func isNetworkExists(cli *client.Client, ctx context.Context, networkName string) (bool, string)  {
	networks, _ := cli.NetworkList(ctx, types.NetworkListOptions{})

	for _, network := range networks {
		if network.Name == networkName {
			return true, network.ID
		}
	}


	return false, ""
}