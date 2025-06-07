package docker

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/docker/docker/api/types/image"
)

// PullImage pulls a Docker image
func (dm *DockerManager) PullImage(ctx context.Context, imageName string) error {
	out, err := dm.cli.ImagePull(ctx, imageName, image.PullOptions{})
	if err != nil {
		return fmt.Errorf("failed to pull image: %v", err)
	}
	defer out.Close()

	_, err = io.Copy(os.Stdout, out)
	return err
}
