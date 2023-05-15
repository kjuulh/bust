package main

import (
	"context"
	"log"

	"git.front.kjuulh.io/kjuulh/bust/pkg/builder"
	"git.front.kjuulh.io/kjuulh/bust/pkg/cli"
	"git.front.kjuulh.io/kjuulh/bust/pkg/pipelines"
)

func main() {
	log.Printf("building bust")

	err := cli.NewCustomGoBuild("golangbin", func(ctx context.Context) error {
		builder, err := builder.New(ctx)
		if err != nil {
			return err
		}

		err = pipelines.
			New(builder).
			WithGolangBin(&pipelines.GolangBinOpts{
				DockerImageOpt: &pipelines.DockerImageOpt{
					ImageName: "bust",
				},
				BuildPath:  "main.go",
				BinName:    "bust",
				BaseImage:  "harbor.front.kjuulh.io/docker-proxy/library/docker:dind",
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
