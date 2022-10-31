package cli

import (
	"git.front.kjuulh.io/kjuulh/bust/pkg/cli/templatecmd"
	"github.com/spf13/cobra"
)

func NewCli() *cobra.Command {
	cmd := &cobra.Command{
		Use: "bust",
	}

	cmd.AddCommand(Build())
	cmd.AddCommand(templatecmd.NewTemplateCmd())

	return cmd
}
