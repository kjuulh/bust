package templatecmd

import (
	"embed"
	"errors"
	"os"
	"text/template"

	"github.com/spf13/cobra"
)

//go:embed templates/gobin_default/*
var gobinDefault embed.FS

func NewInitCmd() *cobra.Command {
	var (
		template string
		name     string
	)

	cmd := &cobra.Command{
		Use: "init",
		RunE: func(cmd *cobra.Command, args []string) error {
			if err := cmd.ParseFlags(args); err != nil {
				return err
			}

			switch template {
			case "gobin_default":
				if err := initializeTemplate(&gobinDefault, name); err != nil {
					return err
				}
				break
			default:
				return errors.New("could not find matching templates, please run [bust template ls] instead")
			}

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&template, "template", "p", "", "The template to initialize")
	cmd.MarkPersistentFlagRequired("template")

	cmd.PersistentFlags().StringVarP(&name, "name", "n", "", "The name into the template")
	cmd.MarkPersistentFlagRequired("name")

	return cmd
}

func initializeTemplate(t *embed.FS, name string) error {
	tinit := template.
		Must(
			template.
				New("").
				Delims("[[", "]]").
				ParseFS(t, "templates/gobin_default/*"),
		)
	type data struct {
		Name string
	}

	droneWriter, err := os.Create(".drone.yml")
	if err != nil {
		return err
	}

	return tinit.ExecuteTemplate(droneWriter, ".drone.yml", data{
		Name: name,
	})
}
