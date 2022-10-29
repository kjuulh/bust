package tasks

import (
	"context"
	"log"
	"os"

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

					log.Println("listing files in /src/")
					dir, err = os.ReadDir("/src/")
					if err == nil {
						for _, d := range dir {
							log.Printf("content: %s\n", d.Name())
						}
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
