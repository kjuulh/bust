package pipelines

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"time"

	"dagger.io/dagger"
	"git.front.kjuulh.io/kjuulh/byg"
	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/tasks/container"
	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/tasks/golang"
	golangbin "git.front.kjuulh.io/kjuulh/dagger-go/pkg/tasks/golang-bin"
)

type DockerImageOpt struct {
	ImageName string
	ImageTag  string
}

type GolangBinOpts struct {
	*DockerImageOpt
	BuildPath string
	BinName   string
}

func (p *Pipeline) WithGolangBin(opts *GolangBinOpts) *Pipeline {
	log.Printf("building image: %s", opts.ImageName)

	client := p.builder.Dagger
	ctx := context.Background()

	var (
		bin        dagger.FileID
		build      *dagger.Container
		scratch    *dagger.Container
		finalImage *dagger.Container
	)

	pipeline := byg.
		New().
		Step(
			"build golang",
			byg.Step{
				Execute: func(_ byg.Context) error {
					var err error
					c := container.LoadImage(client, "harbor.front.kjuulh.io/docker-proxy/library/golang")
					c, err = container.MountCurrent(ctx, client, c, "/src")
					if err != nil {
						return err
					}
					c = container.Workdir(c, "/src")

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
			byg.Step{
				Execute: func(ctx byg.Context) error {
					scratch = container.LoadImage(client, "harbor.front.kjuulh.io/docker-proxy/library/busybox")
					return nil
				},
			},
		).
		Step(
			"create-production-image",
			byg.Step{
				Execute: func(ctx byg.Context) error {
					tempmount := fmt.Sprintf("/tmp/%s", opts.BinName)
					usrbin := fmt.Sprintf("/usr/bin/%s", opts.BinName)
					c := container.MountFileFromLoaded(scratch, bin, tempmount)
					c = c.Exec(dagger.ContainerExecOpts{
						Args: []string{"mkdir", "-p", "/usr/bin"},
					})
					c = c.Exec(dagger.ContainerExecOpts{
						Args: []string{"cp", tempmount, usrbin},
					})
					finalImage = c.WithEntrypoint([]string{opts.BinName})

					return nil
				},
			},
			byg.Step{
				Execute: func(_ byg.Context) error {
					return golang.Test(ctx, build)
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

					tag := fmt.Sprintf("harbor.front.kjuulh.io/kjuulh/%s:%s", opts.ImageName, opts.ImageTag)

					_, err := finalImage.Publish(ctx, tag)
					return err
				},
			},
		)

	p.add(pipeline)

	return p
}