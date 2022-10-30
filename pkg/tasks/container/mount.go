package container

import (
	"context"
	"log"

	"dagger.io/dagger"
)

func MountCurrent(ctx context.Context, client *dagger.Client, container *dagger.Container, into string) (*dagger.Container, error) {
	log.Printf("mounting current working directory into path: %s", into)
	src, err := client.
		Host().
		Workdir().
		Read().
		ID(ctx)
	if err != nil {
		return nil, err
	}

	return container.WithMountedDirectory(into, src), nil
}

func MountFileFromLoaded(container *dagger.Container, bin dagger.FileID, path string) *dagger.Container {
	log.Printf("mounting binary into container: into (path=%s)", path)
	return container.WithMountedFile(path, bin)
}
