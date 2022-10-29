package tasks

import (
	"context"
	"log"

	"dagger.io/dagger"
	"git.front.kjuulh.io/kjuulh/byg"
	"git.front.kjuulh.io/kjuulh/dagger-go/internal"
)

func Build(builder *internal.Builder, imageTag string) error {
	log.Printf("building image: %s", imageTag)

	client := builder.Dagger
	ctx := context.Background()

	return byg.
		New().
		Step(
			"build golang",
			byg.Step{
				Execute: func(_ byg.Context) error {
					src, err := client.
						Host().
						Workdir().
						Read().
						ID(context.Background())
					if err != nil {
						return err
					}

					golang := client.Container().From("golang:latest")
					golang = golang.WithMountedDirectory("/src", src).WithWorkdir("/src")
					_, err = golang.Exec(dagger.ContainerExecOpts{
						Args: []string{"go", "build", "-o", "build/"},
					}).ExitCode(ctx)

					return err
				},
			}).
		Execute(context.Background())

}
