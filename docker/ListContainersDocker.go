package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
)

// ListContainers lists all containers
func (dm *DockerManager) ListContainers(ctx context.Context, all bool) ([]container.Summary, error) {
	containers, err := dm.cli.ContainerList(ctx, container.ListOptions{All: all})
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %v", err)
	}
	return containers, nil
}
