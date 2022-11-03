package cmd

import (
	"github.com/spf13/cobra"
	"myclass_service/src/config"
)

type command = func(config config.IConfig) *cobra.Command

func GetRoot(config config.IConfig) *cobra.Command {
	cmds := []command{server, buildErr, genError}

	root := &cobra.Command{}

	for _, item := range cmds {
		root.AddCommand(item(config))
	}

	return root
}
