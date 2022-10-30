package golang

import (
	"context"
	"log"

	"dagger.io/dagger"
)

func Test(ctx context.Context, container *dagger.Container) error {
	log.Printf("testing: image")
	c := container.Exec(dagger.ContainerExecOpts{
		Args: []string{"go", "test", "./..."},
	})

	_, err := c.ExitCode(ctx)
	if err != nil {
		return err
	}

	return nil
}
