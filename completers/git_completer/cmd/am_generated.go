package cmd

import (
	"github.com/spf13/cobra"
)

var amCmd = &cobra.Command{
	Use:   "am",
	Short: "Apply a series of patches from a mailbox",
	Run: func(cmd *cobra.Command, args []string) {
	},
}

func init() {
	amCmd.Flags().BoolP("3way", "3", false, "allow fall back on 3way merging if needed")
	amCmd.Flags().Bool("abort", false, "restore the original branch and abort the patching operation.")
	amCmd.Flags().StringP("C", "C", "", "pass it through git-apply")
	amCmd.Flags().Bool("committer-date-is-author-date", false, "lie about committer date")
	amCmd.Flags().Bool("continue", false, "continue applying patches after resolving a conflict")
	amCmd.Flags().BoolP("scissors", "c", false, "strip everything before a scissors line")
	amCmd.Flags().String("directory", "", "pass it through git-apply")
	amCmd.Flags().String("exclude", "", "pass it through git-apply")
	amCmd.Flags().Bool("ignore-date", false, "use current timestamp for author date")
	amCmd.Flags().Bool("ignore-space-change", false, "pass it through git-apply")
	amCmd.Flags().Bool("ignore-whitespace", false, "pass it through git-apply")
	amCmd.Flags().BoolP("interactive", "i", false, "run interactively")
	amCmd.Flags().String("include", "", "pass it through git-apply")
	amCmd.Flags().Bool("keep-cr", false, "pass --keep-cr flag to git-mailsplit for mbox format")
	amCmd.Flags().Bool("keep-non-patch", false, "pass -b flag to git-mailinfo")
	amCmd.Flags().BoolP("keep", "k", false, "pass -k flag to git-mailinfo")
	amCmd.Flags().BoolP("message-id", "m", false, "pass -m flag to git-mailinfo")
	amCmd.Flags().Bool("no-keep-cr", false, "do not pass --keep-cr flag to git-mailsplit independent of am.keepcr")
	amCmd.Flags().String("patch-format", "", "format the patch(es) are in")
	amCmd.Flags().StringP("p", "p", "", "pass it through git-apply")
	amCmd.Flags().BoolP("quiet", "q", false, "be quiet")
	amCmd.Flags().Bool("quit", false, "abort the patching operation but keep HEAD where it is.")
	amCmd.Flags().Bool("reject", false, "pass it through git-apply")
	amCmd.Flags().Bool("rerere-autoupdate", false, "update the index with reused conflict resolution if possible")
	amCmd.Flags().String("resolvemsg", "", "override error message when patch failure occurs")
	amCmd.Flags().BoolP("resolved", "r", false, "synonyms for --continue")
	amCmd.Flags().StringP("gpg-sign", "S", "", "GPG-sign commits")
	amCmd.Flags().String("show-current-patch", "", "show the patch being applied")
	amCmd.Flags().Bool("skip", false, "skip the current patch")
	amCmd.Flags().BoolP("signoff", "s", false, "add a Signed-off-by line to the commit message")
	amCmd.Flags().BoolP("utf8", "u", false, "recode into utf8 (default)")
	amCmd.Flags().String("whitespace", "", "pass it through git-apply")
	rootCmd.AddCommand(amCmd)
}
