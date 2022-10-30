package cli

import (
	"errors"
	"fmt"
	"os"

	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/builder"
	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/pipelines"
	"github.com/spf13/cobra"
)

func BuildGolangBin() *cobra.Command {
	cmd := &cobra.Command{
		Use: "golangbin",
		RunE: func(cmd *cobra.Command, args []string) error {
			repoName := os.Getenv("DRONE_REPO_NAME")
			if repoName == "" {
				return errors.New("could not find DRONE_REPO_NAME")
			}
			imageTag := fmt.Sprintf("harbor.front.kjuulh.io/library/%s", repoName)

			ctx := cmd.Context()

			builder, err := builder.New(ctx)
			if err != nil {
				return err
			}
			defer builder.CleanUp()

			return pipelines.
				New(builder).
				WithGolangBin(&pipelines.GolangBinOpts{
					DockerImageOpt: &pipelines.DockerImageOpt{
						ImageName: imageTag,
					},
					BuildPath: "main.go",
					BinName:   "main",
				}).
				Execute(ctx)
		},
	}

	return cmd
}
