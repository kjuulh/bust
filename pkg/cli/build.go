package cli

import (
	"github.com/spf13/cobra"
)

func Build() *cobra.Command {
	cmd := &cobra.Command{
		Use: "build",
	}

	cmd.AddCommand(
		BuildGolangBin(),
		BuildDocker(),
	)

	return cmd
}
