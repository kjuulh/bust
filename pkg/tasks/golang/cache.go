package golang

import (
	"context"

	"dagger.io/dagger"
)

func Cache(ctx context.Context, client *dagger.Client, container *dagger.Container) (*dagger.Container, error) {
	cacheKey := "gomods"
	cacheID, err := client.CacheVolume(cacheKey).ID(ctx)
	if err != nil {
		return nil, err
	}
	return container.
		WithMountedCache(cacheID, "/cache").
		WithEnvVariable("GOMODCACHE", "/cache"), nil
}
