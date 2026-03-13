package docker

import (
	"context"
	"os"
	"os/exec"

	"github.com/moby/moby/client"
)

type DockerAPI struct {
}

func NewDockerAPI() *DockerAPI {

	ctx := context.Background()
	cli, err := client.New(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		panic("Docker is not installed or not running")
	}
	_, err = cli.Ping(ctx, client.PingOptions{})
	if err != nil {
		panic("Docker is not running")
	}

	return &DockerAPI{}
}

func (d *DockerAPI) GetContainerList() ([]string, error) {
	ctx := context.Background()
	cli, err := client.New(client.FromEnv, client.WithAPIVersionNegotiation())

	if err != nil {
		return nil, err
	}

	containers, err := cli.ContainerList(ctx, client.ContainerListOptions{})
	if err != nil {
		return nil, err
	}

	var containerNames []string
	for _, container := range containers.Items {
		containerNames = append(containerNames, container.Names[0])
	}

	return containerNames, nil
}

func (d *DockerAPI) ComposeContainer(composeFilePath string) error {
	cmd := exec.Command("docker-compose", "-f", composeFilePath, "up", "-d")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}
