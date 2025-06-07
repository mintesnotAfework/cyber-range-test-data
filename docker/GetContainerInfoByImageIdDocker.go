package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/filters"
)

// GetContainersByImage returns all containers using a specific image (running or stopped)
func (dm *DockerManager) GetContainersByImage(ctx context.Context, imageNameOrID string) ([]container.Summary, error) {
	// Setup filter to find containers by ancestor image
	filter := filters.NewArgs()
	filter.Add("ancestor", imageNameOrID)

	// List containers with the filter (include stopped containers)
	containers, err := dm.cli.ContainerList(ctx, container.ListOptions{
		All:     true,
		Filters: filter,
	})
	if err != nil {
		return nil, fmt.Errorf("failed to list containers: %v", err)
	}

	return containers, nil
}
