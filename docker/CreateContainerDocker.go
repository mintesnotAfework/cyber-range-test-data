package docker

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/network"
	"github.com/joho/godotenv"
)

// StartContainer starts a container from an image
func (dm *DockerManager) CreateContainer(ctx context.Context, imageName string, containerName string, flag string) (string, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return "", errors.New("can not find the .env file")
	}

	bridge_name := os.Getenv("BRIDGED_CUSTOM_NETWORK")
	resp, err := dm.cli.ContainerCreate(
		ctx,
		&container.Config{
			Image: imageName,
			Tty:   false,
			Env: []string{
				"ME_CTF_FLAG=" + flag,
			},
		},
		&container.HostConfig{},
		&network.NetworkingConfig{
			EndpointsConfig: map[string]*network.EndpointSettings{
				bridge_name: {}, // Attach here
			},
		}, nil, containerName,
	)
	if err != nil {
		return "", fmt.Errorf("failed to create container: %v", err)
	}

	return resp.ID, nil
}
