package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

// RemoveContainer removes a container
func (dm *DockerManager) RemoveContainer(ctx context.Context, containerID string) error {
	if err := dm.cli.ContainerRemove(ctx, containerID, container.RemoveOptions{}); err != nil {
		return fmt.Errorf("failed to remove container: %v", err)
	}
	return nil
}
