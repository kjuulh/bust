package templatecmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var templates = []string{"docker", "gobin_default", "default", "rustbin_default"}

func NewLsCmd() *cobra.Command {
	return &cobra.Command{
		Use: "ls",
		RunE: func(cmd *cobra.Command, args []string) error {

			for _, t := range templates {
				fmt.Printf("%s\n", t)
			}

			return nil
		},
	}
}
