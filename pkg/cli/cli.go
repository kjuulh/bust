package cli

import (
	"log"

	"git.front.kjuulh.io/kjuulh/dagger-go/internal"
	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/tasks"
	"github.com/spf13/cobra"
)

func Build() *cobra.Command {
	var (
		imageTag string
	)

	cmd := &cobra.Command{
		Use: "build",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ParseFlags(args); err != nil {
				return err
			}

			ctx := cmd.Context()

			log.Printf("Building image: %s\n", imageTag)

			client, err := internal.New(ctx)
			if err != nil {
				return err
			}
			defer client.CleanUp()

			return tasks.Build(client, imageTag)
		},
	}

	cmd.PersistentFlags().StringVar(&imageTag, "image-tag", "", "the url for which to tag the docker image, defaults to private url, with repo as image name")
	cmd.MarkPersistentFlagRequired("image-tag")

	return cmd
}
