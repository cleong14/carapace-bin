package cmd

import (
	"github.com/spf13/cobra"
)

var revertCmd = &cobra.Command{
	Use:   "revert",
	Short: "Revert some existing commits",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	revertCmd.Flags().Bool("abort", false, "cancel revert or cherry-pick sequence")
	revertCmd.Flags().String("cleanup", "", "how to strip spaces and #comments from message")
	revertCmd.Flags().Bool("continue", false, "resume revert or cherry-pick sequence")
	revertCmd.Flags().BoolP("edit", "e", false, "edit the commit message")
	revertCmd.Flags().BoolP("mainline", "m", false, "<parent-number>    select mainline parent")
	revertCmd.Flags().BoolP("no-commit", "n", false, "don't automatically commit")
	revertCmd.Flags().Bool("quit", false, "end revert or cherry-pick sequence")
	revertCmd.Flags().Bool("rerere-autoupdate", false, "update the index with reused conflict resolution if possible")
	revertCmd.Flags().StringP("gpg-sign", "S", "", "GPG sign commit")
	revertCmd.Flags().Bool("skip", false, "skip current commit and continue")
	revertCmd.Flags().BoolP("signoff", "s", false, "add Signed-off-by:")
	revertCmd.Flags().String("strategy", "", "merge strategy")
	revertCmd.Flags().BoolP("strategy-option", "X", false, "<option>    option for merge strategy")
	rootCmd.AddCommand(revertCmd)
}
