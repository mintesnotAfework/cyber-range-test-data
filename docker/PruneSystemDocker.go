package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/api/types/image"
)

// PruneSystem removes unused Docker objects
func (dm *DockerManager) PruneSystem(ctx context.Context) (*image.PruneReport, error) {
	report, err := dm.cli.ImagesPrune(ctx, filters.NewArgs())
	if err != nil {
		return nil, fmt.Errorf("failed to prune system: %v", err)
	}
	return &report, nil
}
