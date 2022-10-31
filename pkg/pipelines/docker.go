package pipelines

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"dagger.io/dagger"
	"git.front.kjuulh.io/kjuulh/byg"
)

type DockerOpt struct {
	*DockerImageOpt
	Path string
}

func (p *Pipeline) WithDocker(opts *DockerOpt) *Pipeline {
	log.Printf("building image: %s", opts.ImageName)

	client := p.builder.Dagger
	ctx := context.Background()

	var (
		finalImage *dagger.Container
	)

	pipeline := byg.
		New().
		Step(
			"build image",
			byg.Step{
				Execute: func(_ byg.Context) error {
					var err error

					dir, err := client.Host().Workdir().Read().ID(ctx)
					if err != nil {
						return err
					}
					finalImage = client.Container().Build(dir, dagger.ContainerBuildOpts{Dockerfile: opts.Path})
					if _, err = finalImage.ExitCode(ctx); err != nil {
						return err
					}

					return nil
				},
			},
		).
		Step(
			"upload-image",
			byg.Step{
				Execute: func(_ byg.Context) error {

					if opts.ImageTag == "" {
						opts.ImageTag = strconv.FormatInt(time.Now().UTC().UnixMilli(), 10)
					}

					tag := fmt.Sprintf("harbor.server.kjuulh.io/kjuulh/%s:%s", opts.ImageName, opts.ImageTag)

					_, err := finalImage.Publish(ctx, tag)
					return err
				},
			},
		)

	p.add(pipeline)

	return p
}
