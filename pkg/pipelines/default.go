package pipelines

import (
	"context"
	"log"

	"git.front.kjuulh.io/kjuulh/byg"
)

func (p *Pipeline) WithDefault() error {
	return byg.
		New().
		Step(
			"default step",
			byg.Step{
				Execute: func(ctx byg.Context) error {
					log.Println("Hello, world!")
					return nil
				},
			}).
		Execute(context.Background())

}
