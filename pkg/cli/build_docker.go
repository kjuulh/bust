package cli

import (
	"errors"
	"os"

	"git.front.kjuulh.io/kjuulh/bust/pkg/builder"
	"git.front.kjuulh.io/kjuulh/bust/pkg/pipelines"
	"github.com/spf13/cobra"
)

func BuildDocker() *cobra.Command {
	cmd := &cobra.Command{
		Use: "docker",
		RunE: func(cmd *cobra.Command, args []string) error {
			repoName := os.Getenv("DRONE_REPO_NAME")
			if repoName == "" {
				return errors.New("could not find DRONE_REPO_NAME")
			}

			ctx := cmd.Context()

			builder, err := builder.New(ctx)
			if err != nil {
				return err
			}
			defer builder.CleanUp()

			return pipelines.
				New(builder).
				WithDocker(&pipelines.DockerOpt{
					DockerImageOpt: &pipelines.DockerImageOpt{ImageName: repoName},
					Path:           "Dockerfile",
				}).
				Execute(ctx)
		},
	}

	return cmd
}
