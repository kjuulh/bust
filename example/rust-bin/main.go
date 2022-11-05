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
		WithRustBin(&pipelines.RustBinOpts{
			DockerImageOpt: &pipelines.DockerImageOpt{
				ImageName: "rust-bin",
			},
			BinName: "rust-bin",
		}).
		Execute(ctx)
}
