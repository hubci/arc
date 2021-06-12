package cmd

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var followCmd = &cobra.Command{
	Use:   "follow",
	Short: "Follow this repository on CircleCI",
	RunE: func(cmd *cobra.Command, args []string) error {

		if os.Getenv("CIRCLECI_TOKEN") == "" {
			return fmt.Errorf("The CircleCI token hasn't been set in the evironment.")
		}

		output, err := exec.Command("git", "remote", "get-url", "origin").CombinedOutput()
		if err != nil {
			fmt.Printf(err.Error())
		}

		remoteURL := strings.Split(string(output), ":")[1]
		org := strings.Split(remoteURL, "/")[0]
		repo := strings.Split(remoteURL, "/")[1]
		repo = repo[0 : len(repo)-5]

		reqURL := fmt.Sprintf("https://circleci.com/api/v1.1/project/gh/%s/%s/follow?circle-token=%s", org, repo, os.Getenv("CIRCLECI_TOKEN"))

		resp, err := http.Post(reqURL, "application/json", nil)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := io.ReadAll(resp.Body)
		fmt.Printf("%s", body)

		return nil
	},
}

func init() {

	rootCmd.AddCommand(followCmd)
}
