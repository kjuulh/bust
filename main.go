package daggergo

import (
	"git.front.kjuulh.io/kjuulh/dagger-go/cmd"
	"github.com/spf13/cobra"
)

func CreateCmd() *cobra.Command {
	cobracmd := &cobra.Command{
		Use: "dagger-go",
	}

	cobracmd.AddCommand(cmd.Build())

	return cobracmd
}
