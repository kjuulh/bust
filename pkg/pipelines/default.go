package pipelines

import (
	"log"

	"git.front.kjuulh.io/kjuulh/byg"
)

func (p *Pipeline) WithDefault() *byg.Builder {
	return byg.
		New().
		Step(
			"default step",
			byg.Step{
				Execute: func(ctx byg.Context) error {
					log.Println("Hello, world!")
					return nil
				},
			})
}
