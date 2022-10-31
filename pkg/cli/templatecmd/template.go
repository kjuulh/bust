package templatecmd

import "github.com/spf13/cobra"

func NewTemplateCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use: "template",
	}

	cmd.AddCommand(NewInitCmd())

	return cmd
}
