package cmd

import (
	"errors"
	"fmt"
	"os/exec"
	"runtime"
	"strings"

	"github.com/spf13/cobra"
)

// flags
var github, circleci bool

var browseCmd = &cobra.Command{
	Use:   "browse",
	Short: "Open this repository in the browser",
	RunE: func(cmd *cobra.Command, args []string) error {

		rawRemotes, err := exec.Command("git", "remote").CombinedOutput()
		if err != nil {
			return errors.New("Failed to retrieve git branch names.")
		}

		remotes := strings.Split(string(rawRemotes), "\n")

		var desiredRemote string
		for _, remote := range remotes {
			if remote == "origin" {
				desiredRemote = remote
			}
		}

		if desiredRemote == "" {
			desiredRemote = remotes[0]
		}

		output, err := exec.Command("git", "remote", "get-url", desiredRemote).CombinedOutput()
		if err != nil {
			fmt.Printf(err.Error())
		}

		remoteURL := string(output)

		remoteURL = strings.TrimPrefix(remoteURL, "git@github.com:")
		remoteURL = strings.TrimPrefix(remoteURL, "https://github.com/")
		remoteURL = strings.TrimPrefix(remoteURL, "http://github.com/")
		remoteURL = strings.TrimSuffix(remoteURL, "\n")
		remoteURL = strings.TrimSuffix(remoteURL, ".git")

		org := strings.Split(remoteURL, "/")[0]
		repo := strings.Split(remoteURL, "/")[1]

		if circleci {

			browserURL := fmt.Sprintf("https://app.circleci.com/pipelines/github/%s/%s", org, repo)
			openInBrowser(browserURL)
		}

		if github {

			browserURL := fmt.Sprintf("https://github.com/%s/%s", org, repo)
			openInBrowser(browserURL)
		}

		return nil
	},
}

func init() {

	browseCmd.Flags().BoolVarP(&github, "github", "g", true, "open in GitHub")
	browseCmd.Flags().BoolVarP(&circleci, "circleci", "c", false, "open in CircleCI")
	rootCmd.AddCommand(browseCmd)
}

func openInBrowser(browserURL string) {

	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", browserURL).Start()
		if err != nil {
			fmt.Printf(err.Error())
		}
	case "darwin":
		err = exec.Command("open", browserURL).Start()
		if err != nil {
			fmt.Printf(err.Error())
		}
	default:
		fmt.Println("Error: unsupported OS.")
	}
	if err != nil {
		fmt.Println("Failed to open repo in browser.")
	}
}
