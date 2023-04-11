package typed

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

type Container struct {
	Name string
}

var dockerd *client.Client = nil

func initDockerClient() {
	if dockerd == nil {
		c, err := client.NewEnvClient()
		if err != nil {
			panic(err)
		}
		serverVersion, err := c.ServerVersion(context.Background())
		if err != nil {
			fmt.Printf("failed to get docker server version, err:%+v\n", err)
		}
		fmt.Printf("docker verion: %s, api version:%s\n", serverVersion.Version, c.ClientVersion())
		dockerd = c
	}
}

func listContainer() []types.Container {
	containers, err := dockerd.ContainerList(context.Background(), types.ContainerListOptions{})
	if err != nil {
		fmt.Printf("failed to list container, err:%+v\n", err)
	}
	return containers
}

func subscribeContainer() {
	//dockerd
}

func enrichContainer() {

}
