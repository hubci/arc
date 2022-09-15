package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
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
			return fmt.Errorf("The CircleCI token hasn't been set in the environment.")
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

		req, err := http.NewRequest("POST", reqURL, nil)
		if err != nil {
			return err
		}
		req.Header.Set("Accept", "application/json")

		resp, err := http.DefaultClient.Do(req)
		if err != nil {
			return err
		}
		defer resp.Body.Close()

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}

		var followResp map[string]interface{}

		err = json.Unmarshal(body, &followResp)
		if err != nil {
			return errors.New("Failed to unmarshal CircleCI's API reponse.")
		}

		if followResp["message"] != nil {

			msg := followResp["message"].(string)

			if strings.Contains(msg, "Not Found") {
				fmt.Println("Error: Failed to follow project as it wasn't found.")
			} else if strings.Contains(msg, "For security purposes") {
				fmt.Println("Error: You don't have permission to follow project.")
			} else {
				fmt.Printf("Error: Failed to follow project. Here's what CircleCI's API returns: %s", msg)
			}

			return nil

		}

		fmt.Println("The project has been followed successfully.")

		return nil
	},
}

func init() {

	rootCmd.AddCommand(followCmd)
}
