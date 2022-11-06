package pipelines

import (
	"context"
	"fmt"
	"log"
	"path"
	"strconv"
	"time"

	"dagger.io/dagger"
	"git.front.kjuulh.io/kjuulh/bust/pkg/tasks/container"
	rustbin "git.front.kjuulh.io/kjuulh/bust/pkg/tasks/rust-bin"
	"git.front.kjuulh.io/kjuulh/byg"
)

type RustBinOpts struct {
	*DockerImageOpt
	BinName   string
	BaseImage string
}

func (p *Pipeline) WithRustBin(opts *RustBinOpts) *Pipeline {
	log.Printf("building image: %s", opts.ImageName)

	client := p.builder.Dagger
	ctx := context.Background()

	var (
		bin        dagger.FileID
		finalImage *dagger.Container
	)

	pipeline := byg.
		New().
		Step(
			"build rust",
			byg.Step{
				Execute: func(_ byg.Context) error {
					var err error
					c := container.LoadImage(client, "harbor.server.kjuulh.io/docker-proxy/library/rust:buster")
					c = c.Exec(dagger.ContainerExecOpts{
						Args: []string{
							"apt", "update", "-y",
						},
					})
					if _, err := c.ExitCode(ctx); err != nil {
						return err
					}
					c = c.Exec(dagger.ContainerExecOpts{
						Args: []string{
							"apt", "install", "musl-tools", "-y",
						},
					})
					if _, err := c.ExitCode(ctx); err != nil {
						return err
					}

					c = c.Exec(dagger.ContainerExecOpts{
						Args: []string{
							"rustup",
							"target",
							"add",
							"x86_64-unknown-linux-musl",
						},
					})
					if _, err := c.ExitCode(ctx); err != nil {
						return err
					}

					c, err = container.MountCurrent(ctx, client, c, "/src")
					if err != nil {
						return err
					}
					c = container.Workdir(c, "/src")

					if bin, err = rustbin.Build(ctx, c, opts.BinName); err != nil {
						return err
					}

					return err
				},
			},
		).
		Step(
			"create-production-image",
			byg.Step{
				Execute: func(_ byg.Context) error {
					if opts.BaseImage == "" {
						opts.BaseImage = "harbor.server.kjuulh.io/docker-proxy/library/alpine"
					}

					binpath := "/usr/bin"
					usrbin := path.Join(binpath, opts.BinName)
					c := container.LoadImage(client, opts.BaseImage)
					c = c.Exec(dagger.ContainerExecOpts{
						Args: []string{"mkdir", "-p", binpath},
					})
					_, err := c.ExitCode(ctx)
					if err != nil {
						return err
					}

					c, err = container.MountFileFromLoaded(ctx, c, bin, usrbin)
					if err != nil {
						return err
					}
					finalImage = c

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
