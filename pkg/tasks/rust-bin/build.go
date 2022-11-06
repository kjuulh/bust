package rustbin

import (
	"context"
	"fmt"
	"log"

	"dagger.io/dagger"
)

func Build(ctx context.Context, container *dagger.Container, binName string) (dagger.FileID, error) {
	log.Printf("building binary: (binName=%s)", binName)
	c := container.Exec(dagger.ContainerExecOpts{
		Args: []string{
			"cargo",
			"build",
			"--release",
			"--target=x86_64-unknown-linux-musl",
		},
	})
	if _, err := c.ExitCode(ctx); err != nil {
		return "", err
	}

	bin, err := c.File(fmt.Sprintf("target/x86_64-unknown-linux-musl/release/%s", binName)).ID(ctx)
	if err != nil {
		return "", err
	}

	return bin, nil
}
