package pipelines

import (
	"context"

	"git.front.kjuulh.io/kjuulh/byg"
	"git.front.kjuulh.io/kjuulh/dagger-go/internal"
	"golang.org/x/sync/errgroup"
)

type Pipeline struct {
	builder   *internal.Builder
	pipelines []*byg.Builder
}

func New(builder *internal.Builder) *Pipeline {
	return &Pipeline{builder: builder}
}

func (p *Pipeline) WithCustom(custom func(p *Pipeline) *byg.Builder) {
	p.add(custom(p))
}

func (p *Pipeline) Execute(ctx context.Context) error {
	errgroup, ctx := errgroup.WithContext(ctx)

	for _, pipeline := range p.pipelines {
		pipeline := pipeline // Allocate for closure

		errgroup.Go(func() error {
			return pipeline.Execute(ctx)
		})
	}

	if err := errgroup.Wait(); err != nil {
		return err
	}

	return nil
}

func (p *Pipeline) add(pipeline *byg.Builder) {
	p.pipelines = append(p.pipelines, pipeline)
}
