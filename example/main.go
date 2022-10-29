package main

import (
	"context"
	"log"

	"git.front.kjuulh.io/kjuulh/dagger-go/internal"
	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/tasks"
)

func main() {
	ctx := context.Background()

	if err := run(ctx); err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	builder, err := internal.New(ctx)
	if err != nil {
		return err
	}
	defer builder.CleanUp()

	return tasks.Build(builder, "some-image")
}
