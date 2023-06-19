package cmd

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/repowarden/cli/warden/vcsurl"
	"github.com/spf13/cobra"
)

var (
	cloneCmd = &cobra.Command{
		Use:   "clone",
		Short: "Clones a git repo to the ~/Repos/org/name structure",
		RunE: func(cmd *cobra.Command, args []string) error {

			repo, err := vcsurl.Parse(args[0])
			if err != nil {
				return err
			}

			homeDir, err := os.UserHomeDir()
			if err != nil {
				return err
			}

			// create directories if needed
			filepath := homeDir + "/Repos/" + repo.Owner

			err = os.MkdirAll(filepath, 0775)
			if err != nil && !os.IsExist(err) {
				return err
			}

			aCmd := exec.Command("git", "clone", "--recurse-submodules", args[0], filepath+"/"+repo.Name)
			stdout, err := aCmd.Output()
			if err != nil {
				fmt.Printf("%s/n", stdout)
				return err
			}

			fmt.Printf("Cloned successfully to %s\n", filepath+"/"+repo.Name)

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(cloneCmd)
}
