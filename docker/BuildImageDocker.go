package docker

import (
	"archive/tar"
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/docker/docker/api/types"
)

// BuildImage builds a Docker image from a Dockerfile
func (dm *DockerManager) BuildImage(ctx context.Context, contextDir string, dockerfile string, tag string) error {
	// Create a tar archive of the build context
	buf := new(bytes.Buffer)
	tw := tar.NewWriter(buf)
	defer tw.Close()

	// Walk through the build context directory
	err := filepath.Walk(contextDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// Skip directories and the .dockerignore files
		if info.IsDir() {
			return nil
		}

		// Create tar header
		header, err := tar.FileInfoHeader(info, info.Name())
		if err != nil {
			return err
		}

		// Make paths relative to the build context
		relPath, err := filepath.Rel(contextDir, path)
		if err != nil {
			return err
		}
		header.Name = relPath

		// Write header
		if err := tw.WriteHeader(header); err != nil {
			return err
		}

		// Write file content
		if !info.IsDir() {
			data, err := os.Open(path)
			if err != nil {
				return err
			}
			defer data.Close()
			if _, err := io.Copy(tw, data); err != nil {
				return err
			}
		}
		return nil
	})

	if err != nil {
		return fmt.Errorf("failed to create build context: %v", err)
	}

	// Important: Close the tar writer before using the buffer
	tw.Close()

	buildOptions := types.ImageBuildOptions{
		Tags:       []string{tag},
		Dockerfile: dockerfile, // Relative path within the context
		Remove:     true,
	}

	// Use the tar archive as build context
	buildResponse, err := dm.cli.ImageBuild(ctx, buf, buildOptions)
	if err != nil {
		return fmt.Errorf("failed to build image: %v", err)
	}
	defer buildResponse.Body.Close()

	// Print build output
	_, err = io.Copy(os.Stdout, buildResponse.Body)
	return err
}
