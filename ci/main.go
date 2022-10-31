package main

import (
	"context"
	"log"

	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/builder"
	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/cli"
	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/pipelines"
)

func main() {
	log.Printf("building dagger-go")

	err := cli.NewCustomGoBuild("golangbin", func(ctx context.Context) error {
		builder, err := builder.New(ctx)
		if err != nil {
			return err
		}

		err = pipelines.
			New(builder).
			WithGolangBin(&pipelines.GolangBinOpts{
				DockerImageOpt: &pipelines.DockerImageOpt{
					ImageName: "dagger-go",
				},
				BuildPath:  "main.go",
				BinName:    "dagger-go",
				BaseImage:  "harbor.server.kjuulh.io/docker-proxy/library/docker:dind",
				CGOEnabled: true,
			}).
			Execute(ctx)

		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Fatal(err)
	}
}
