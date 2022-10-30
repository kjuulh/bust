package cli

import (
	"errors"
	"fmt"
	"log"
	"os"

	"git.front.kjuulh.io/kjuulh/dagger-go/internal"
	"git.front.kjuulh.io/kjuulh/dagger-go/pkg/pipelines"
	"github.com/spf13/cobra"
)

func Build(mainGoPath string, imageTag string) *cobra.Command {
	cmd := &cobra.Command{
		Use: "build",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ParseFlags(args); err != nil {
				return err
			}

			if imageTag == "" {
				repoName := os.Getenv("DRONE_REPO_NAME")
				if repoName == "" {
					return errors.New("could not find DRONE_REPO_NAME")
				}
				imageTag = fmt.Sprintf("harbor.front.kjuulh.io/library/%s", repoName)
			}

			ctx := cmd.Context()

			log.Printf("Building image: %s\n", imageTag)

			builder, err := internal.New(ctx)
			if err != nil {
				return err
			}
			defer builder.CleanUp()

			return pipelines.
				New(builder).
				WithGolangBin(&pipelines.GolangBinOpts{
					DockerImageOpt: &pipelines.DockerImageOpt{
						ImageName: "golang-bin",
					},
					BuildPath: "example/golang-bin/main.go",
					BinName:   "golang-bin",
				}).
				Execute(ctx)

		},
	}

	return cmd
}
