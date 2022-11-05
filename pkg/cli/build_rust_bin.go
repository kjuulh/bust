package cli

import (
	"errors"
	"os"

	"git.front.kjuulh.io/kjuulh/bust/pkg/builder"
	"git.front.kjuulh.io/kjuulh/bust/pkg/pipelines"
	"github.com/spf13/cobra"
)

func BuildRustBin() *cobra.Command {
	var (
		binName string
	)

	cmd := &cobra.Command{
		Use: "rustbin",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ParseFlags(args); err != nil {
				return err
			}

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
				WithRustBin(&pipelines.RustBinOpts{
					DockerImageOpt: &pipelines.DockerImageOpt{
						ImageName: repoName,
					},
					BinName: binName,
				}).
				Execute(ctx)
		},
	}

	cmd.PersistentFlags().StringVar(&binName, "bin-name", "", "bin-name is the binary to build, and what will be present in the output folder")

	return cmd
}
