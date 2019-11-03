package createClient

import "github.com/docker/docker/client"

// CreateClient create client for the docker cli
func CreateClient() (*client.Client, error)  {
	// create new docker client
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	return cli, nil
}
