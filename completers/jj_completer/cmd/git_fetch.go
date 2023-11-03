package cmd

import (
	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/pkg/actions/tools/jj"
	"github.com/spf13/cobra"
)

var git_fetchCmd = &cobra.Command{
	Use:   "fetch",
	Short: "Fetch from a Git remote",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(git_fetchCmd).Standalone()

	git_fetchCmd.Flags().Bool("all-remotes", false, "Fetch from all remotes")
	git_fetchCmd.Flags().StringSlice("branch", []string{}, "Fetch only some of the branches")
	git_fetchCmd.Flags().BoolP("help", "h", false, "Print help (see more with '--help')")
	git_fetchCmd.Flags().StringSlice("remote", []string{"origin"}, "The remote to fetch from (only named remotes are supported, can be repeated)")
	gitCmd.AddCommand(git_fetchCmd)

	carapace.Gen(git_fetchCmd).FlagCompletion(carapace.ActionMap{
		"branch": carapace.ActionCallback(func(c carapace.Context) carapace.Action { // TODO verify
			remotes, err := git_fetchCmd.Flags().GetStringSlice("remote")
			if err != nil {
				return carapace.ActionMessage(err.Error())
			}

			batch := carapace.Batch()
			for _, remote := range remotes {
				batch = append(batch, jj.ActionRemoteBranches(remote))
			}
			return batch.ToA()
		}),
		"remote": jj.ActionRemotes(),
	})
}
