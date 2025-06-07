package docker

import (
	"context"
	"fmt"
	"os"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/pkg/stdcopy"
)

// ContainerLogs retrieves container logs
func (dm *DockerManager) ContainerLogs(ctx context.Context, containerID string) error {
	out, err := dm.cli.ContainerLogs(ctx, containerID, container.LogsOptions{
		ShowStdout: true,
		ShowStderr: true,
		Follow:     false,
		Timestamps: false,
	})
	if err != nil {
		return fmt.Errorf("failed to get container logs: %v", err)
	}
	defer out.Close()

	_, err = stdcopy.StdCopy(os.Stdout, os.Stderr, out)
	return err
}
