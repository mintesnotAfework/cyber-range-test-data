package docker

import (
	"context"
	"fmt"
	"strings"

	"github.com/docker/docker/api/types/container"
	"github.com/docker/docker/api/types/image"
	"github.com/docker/docker/client"
)

func (dm *DockerManager) IsContainerStopped(ctx context.Context, containerNameOrID string) (bool, error) {
	containerInfo, err := dm.cli.ContainerInspect(ctx, containerNameOrID)
	if err != nil {
		if client.IsErrNotFound(err) {
			return false, fmt.Errorf("container %s not found", containerNameOrID)
		}
		return false, fmt.Errorf("failed to inspect container: %v", err)
	}

	return !containerInfo.State.Running, nil
}

// ImageHasContainers checks if an image has any associated containers (running or stopped)
func (dm *DockerManager) ImageHasContainers(ctx context.Context, imageNameOrID string) (bool, error) {
	// List all containers (including stopped ones)
	containers, err := dm.cli.ContainerList(ctx, container.ListOptions{
		All: true,
	})
	if err != nil {
		return false, fmt.Errorf("failed to list containers: %v", err)
	}

	// Check each container's image
	for _, c := range containers {
		if c.Image == imageNameOrID || c.ImageID == imageNameOrID {
			return true, nil
		}
	}

	return false, nil
}

// ImageExists checks if an image exists locally by name or ID
func (dm *DockerManager) ImageExists(ctx context.Context, imageNameOrID string) (bool, error) {
	// List all images
	images, err := dm.cli.ImageList(ctx, image.ListOptions{})
	if err != nil {
		return false, fmt.Errorf("failed to list images: %v", err)
	}

	// Check each image's name, tag, and ID
	for _, img := range images {
		// Check against image ID (full or partial)
		if strings.HasPrefix(img.ID, "sha256:"+imageNameOrID) || strings.HasPrefix(img.ID, imageNameOrID) {
			return true, nil
		}

		// Check against repository tags
		for _, tag := range img.RepoTags {
			if tag == imageNameOrID {
				return true, nil
			}
			// Also check without version tag
			if strings.Split(tag, ":")[0] == imageNameOrID {
				return true, nil
			}
		}
	}

	return false, nil
}
