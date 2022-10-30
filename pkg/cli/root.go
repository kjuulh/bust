package cli

import "github.com/spf13/cobra"

func NewCli() *cobra.Command {
	cmd := &cobra.Command{
		Use: "dagger",
	}

	cmd.AddCommand(Build())

	return cmd
}
