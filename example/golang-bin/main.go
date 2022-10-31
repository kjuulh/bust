package main

import (
	"context"
	"log"

	"git.front.kjuulh.io/kjuulh/bust/pkg/builder"
	"git.front.kjuulh.io/kjuulh/bust/pkg/pipelines"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}
func run(ctx context.Context) error {
	builder, err := builder.New(ctx)
	if err != nil {
		return err
	}
	defer builder.CleanUp()

	return pipelines.
		New(builder).
		WithGolangBin(&pipelines.GolangBinOpts{
			DockerImageOpt: &pipelines.DockerImageOpt{
				ImageName: "golang-bin",
			},
			BuildPath: "example/golang-bin/main.go",
			BinName:   "golang-bin",
		}).
		Execute(ctx)
}
