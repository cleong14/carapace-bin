package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/spf13/cobra"
)

var pluginCmd = &cobra.Command{
	Use:   "plugin",
	Short: "Useful commands relating to Gatsby plugins",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(pluginCmd).Standalone()

	rootCmd.AddCommand(pluginCmd)

	carapace.Gen(pluginCmd).PositionalCompletion(
		carapace.ActionValues("docs", "ls"),
	)

	// TODO complete plugins
}
