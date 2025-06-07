package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

// StartContainer starts a container from an image
func (dm *DockerManager) StartContainer(ctx context.Context, containerName string) error {
	if err := dm.cli.ContainerStart(ctx, containerName, container.StartOptions{}); err != nil {
		return fmt.Errorf("failed to start container: %v", err)
	}

	return nil
}
