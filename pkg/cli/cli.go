package cli

import (
	"context"
	"log"

	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/tasks"
	"github.com/spf13/cobra"
)

func Build(requirementsf func(*cobra.Command), buildf func(ctx context.Context) error) *cobra.Command {
	var (
		imageTag string
	)

	cmd := &cobra.Command{
		Use: "build",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ParseFlags(args); err != nil {
				return err
			}

			if imageTag != "" {
				log.Printf("Building image: %s\n", imageTag)
			}

			if buildf != nil {
				return buildf(cmd.Context())
			}

			return tasks.Build(imageTag)
		},
	}

	cmd.PersistentFlags().StringVar(&imageTag, "image-tag", "", "the url for which to tag the docker image, defaults to private url, with repo as image name")

	requirementsf(cmd)

	return cmd
}
