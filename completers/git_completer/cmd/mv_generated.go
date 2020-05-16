package cmd

import (
	"github.com/spf13/cobra"
)

var mvCmd = &cobra.Command{
	Use:   "mv",
	Short: "Move or rename a file, a directory, or a symlink",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	mvCmd.Flags().BoolP("force", "f", false, "force move/rename even if target exists")
	mvCmd.Flags().BoolP("k", "k", false, "skip move/rename errors")
	mvCmd.Flags().BoolP("dry-run", "n", false, "dry run")
	mvCmd.Flags().BoolP("verbose", "v", false, "be verbose")
	rootCmd.AddCommand(mvCmd)
}
