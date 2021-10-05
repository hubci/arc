package cmd

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var runCICmd = &cobra.Command{
	Use:   "run-ci",
	Short: "Run a CircleCI Pipeline based on current branch.",
	Run: func(cmd *cobra.Command, args []string) {

		output, err := exec.Command("git", "remote", "get-url", "origin").CombinedOutput()
		if err != nil {
			fmt.Printf(err.Error())
		}

		remoteURL := strings.Split(string(output), ":")[1]
		org := strings.Split(remoteURL, "/")[0]
		repo := strings.Split(remoteURL, "/")[1]
		repo = repo[0 : len(repo)-5]

		output2, err := exec.Command("git", "branch", "--show-current").CombinedOutput()
		if err != nil {
			fmt.Printf(err.Error())
		}
		branch := strings.TrimSuffix(string(output2), "\n")

		payload := []byte(`{"branch": "` + branch + `"}`)

		reqURL := fmt.Sprintf("https://circleci.com/api/v2/project/gh/%s/%s/pipeline?circle-token=%s", org, repo, os.Getenv("CIRCLECI_TOKEN"))
		req, err := http.NewRequest("POST", reqURL, bytes.NewBuffer(payload))
		req.Header.Set("Content-Type", "application/json")

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error: request failed.")
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			fmt.Println("Error: reading response failed.")
		}

		fmt.Println(string(body))
	},
}

func init() {

	rootCmd.AddCommand(runCICmd)
}
