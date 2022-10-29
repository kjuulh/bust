package internal

import (
	"context"

	"dagger.io/dagger"
)

func CreateBuilder(ctx context.Context) error {
	client, err := dagger.Connect(ctx, dagger.WithLogOutput(os.Stdout))
	if err != nil {
		return err
	}
}
