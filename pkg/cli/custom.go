package cli

import (
	"context"
	"fmt"

	"github.com/spf13/cobra"
)

func NewCustomGoBuild(command string, runf func(ctx context.Context) error) error {
	cmd := &cobra.Command{
		Use: fmt.Sprintf("bust build %s", command),
		RunE: func(cmd *cobra.Command, args []string) error {
			return runf(cmd.Context())
		},
	}

	return cmd.Execute()
}
