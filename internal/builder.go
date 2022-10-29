package internal

import (
	"context"
	"os"

	"dagger.io/dagger"
)

type Builder struct {
	Dagger *dagger.Client
}

func New(ctx context.Context) (*Builder, error) {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return nil, err
	}

	return &Builder{
		Dagger: client,
	}, nil
}

func (b *Builder) CleanUp() error {
	return b.Dagger.Close()
}
