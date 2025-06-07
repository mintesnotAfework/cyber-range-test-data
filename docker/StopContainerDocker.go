package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

// StopContainer stops a running container
func (dm *DockerManager) StopContainer(ctx context.Context, containerID string) error {
	if err := dm.cli.ContainerStop(ctx, containerID, container.StopOptions{}); err != nil {
		return fmt.Errorf("failed to stop container: %v", err)
	}
	return nil
}
