package pipelines

import (
	"context"
	"fmt"
	"log"
	"path"
	"strconv"
	"time"

	"dagger.io/dagger"
	"git.front.kjuulh.io/kjuulh/byg"
	"git.front.kjuulh.io/kjuulh/bust/pkg/tasks/container"
	"git.front.kjuulh.io/kjuulh/bust/pkg/tasks/golang"
	golangbin "git.front.kjuulh.io/kjuulh/bust/pkg/tasks/golang-bin"
)

type DockerImageOpt struct {
	ImageName string
	ImageTag  string
}

type GolangBinOpts struct {
	*DockerImageOpt
	BuildPath           string
	BinName             string
	BaseImage           string
	ExecuteOnEntrypoint bool
	CGOEnabled          bool
}

func (p *Pipeline) WithGolangBin(opts *GolangBinOpts) *Pipeline {
	log.Printf("building image: %s", opts.ImageName)

	client := p.builder.Dagger
	ctx := context.Background()

	var (
		bin        dagger.FileID
		build      *dagger.Container
		finalImage *dagger.Container
	)

	pipeline := byg.
		New().
		Step(
			"build golang",
			byg.Step{
				Execute: func(_ byg.Context) error {
					var err error
					c := container.LoadImage(client, "harbor.server.kjuulh.io/docker-proxy/library/golang")
					c, err = container.MountCurrent(ctx, client, c, "/src")
					if err != nil {
						return err
					}
					c = container.Workdir(c, "/src")

					if opts.CGOEnabled {
						c = c.WithEnvVariable("CGO_ENABLED", "1")
					} else {
						c = c.WithEnvVariable("CGO_ENABLED", "0")
					}

					build, err = golang.Cache(ctx, client, c)
					if err != nil {
						return err
					}

					bin, err = golangbin.Build(ctx, build, opts.BinName, opts.BuildPath)
					if err != nil {
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
					if opts.ExecuteOnEntrypoint {
						finalImage = c.WithEntrypoint([]string{usrbin})
					} else {
						finalImage = c
					}

					return nil
				},
			},
			//byg.Step{
			//	Execute: func(_ byg.Context) error {
			//		return golang.Test(ctx, build)
			//	},
			//},
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
