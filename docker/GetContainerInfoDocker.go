package docker

import (
	"context"
	"errors"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/network"
	"github.com/joho/godotenv"
)

// GetContainerDetails returns detailed information about a container by its ID
func (dm *DockerManager) GetContainerDetails(ctx context.Context, containerID string) (*network.EndpointSettings, error) {
	err := godotenv.Load(".env")
	if err != nil {
		return nil, errors.New("can not find the .env file")
	}

	bridge_name := os.Getenv("BRIDGED_CUSTOM_NETWORK")
	details, err := dm.cli.ContainerInspect(ctx, containerID)
	if err != nil {
		return nil, err
	}

	// For custom bridge networks, check the network name:
	if net, exists := details.NetworkSettings.Networks[bridge_name]; exists {
		return net, nil // Correct IP here
	} else {
		return nil, fmt.Errorf("container not attached to expected network")
	}
}
