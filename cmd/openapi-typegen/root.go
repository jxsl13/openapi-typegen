package main

import (
	"fmt"

	"github.com/jxsl13/openapi-typegen/cmd/openapi-typegen/config"
	"github.com/spf13/cobra"
)

func NewRootCmd() *cobra.Command {
	appName := AppName()
	rootContext := RootContext{}

	// rootCmd represents the run command
	rootCmd := &cobra.Command{
		Use:   fmt.Sprintf("%s openapi.yaml", appName),
		Short: "generate types for a given OpenAPI 3.0.x specification.",
		RunE:  rootContext.RunE,
	}
	// add pre and postrun
	rootContext.PreRunE(rootCmd)
	rootContext.PostRunE(rootCmd)

	rootCmd.AddCommand(NewCompletionCmd(appName))

	return rootCmd
}

type RootContext struct {
	Config config.Config
}

func (c *RootContext) PreRunE(cmd *cobra.Command) {
	c.Config = config.NewDefaultConfig()
	// parse config
	runParser := config.RegisterFlags(&c.Config, true, cmd)

	cmd.PreRunE = func(cmd *cobra.Command, args []string) error {
		if len(args) > 0 {
			// update default value
			c.Config.OpenAPIFilePath = args[0]
		}
		// overwrite with flags or env
		return runParser()
	}
}

func (c *RootContext) PostRunE(cmd *cobra.Command) {
	cmd.PostRunE = func(cmd *cobra.Command, args []string) error {
		return c.Config.Close()
	}
}

func (c *RootContext) RunE(cmd *cobra.Command, args []string) (err error) {
	return nil
}
