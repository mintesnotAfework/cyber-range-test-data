package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types/image"
)

// RemoveImage removes a Docker image
func (dm *DockerManager) RemoveImage(ctx context.Context, imageID string) error {
	_, err := dm.cli.ImageRemove(ctx, imageID, image.RemoveOptions{})
	if err != nil {
		return fmt.Errorf("failed to remove image: %v", err)
	}
	return nil
}
