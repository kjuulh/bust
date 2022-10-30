package golangbin

import (
	"context"
	"fmt"
	"log"

	"dagger.io/dagger"
)

func Build(ctx context.Context, container *dagger.Container, binname string, buildpath string) (dagger.FileID, error) {
	log.Printf("building binary: (binName=%s) into (buildPath=%s)", binname, buildpath)
	binpath := fmt.Sprintf("dist/%s", binname)
	c := container.Exec(dagger.ContainerExecOpts{
		Args: []string{"go", "build", "-o", binpath, buildpath},
	})

	_, err := c.ExitCode(ctx)
	if err != nil {
		return "", err
	}

	bin, err := c.File(binpath).ID(ctx)
	if err != nil {
		return "", err
	}

	return bin, nil
}
