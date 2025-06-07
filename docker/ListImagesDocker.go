package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
)

// ListImages lists all images
func (dm *DockerManager) ListImages(ctx context.Context) ([]image.Summary, error) {
	images, err := dm.cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return nil, fmt.Errorf("failed to list images: %v", err)
	}
	return images, nil
}
