package cmd

import (
	"fmt"
	"time"

	"github.com/spf13/cobra"
)

type spPage struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	URL       string    `url:"url"`
	TimeZone  string    `json:"time_zone"`
	UpdatedAt time.Time `json:"updated_at"`
}

type spStatus struct {
	Indicator   string `json:"indicator"`
	Description string `json:"description"`
}

type spResponse struct {
	Page   *spPage   `json:"page"`
	Status *spStatus `json:"status"`
}

var statusCmd = &cobra.Command{
	Use:   "status",
	Short: "Gives you the status page status of both CircleCI and GitHub",
	Run: func(cmd *cobra.Command, args []string) {

		var cciResp *spResponse
		var ghResp *spResponse
		cciURL := "https://status.circleci.com/api/v2/status.json"
		ghURL := "https://www.githubstatus.com/api/v2/status.json"

		client := New()

		errCCI := client.getJSON(cciURL, &cciResp)
		errGH := client.getJSON(ghURL, &ghResp)

		if errCCI != nil {
			cciResp.Status.Indicator = "can't connect"
		}

		if errGH != nil {
			ghResp.Status.Indicator = "can't connect"
		}

		var cciTabs = ""
		if cciResp.Status.Indicator == "none" {
			cciResp.Status.Indicator = ""
			cciTabs = "\t"
		}

		var ghTabs = ""
		if ghResp.Status.Indicator == "none" {
			ghResp.Status.Indicator = ""
			ghTabs = "\t"
		}

		fmt.Println("Reporting status page results...")
		fmt.Println("")
		fmt.Printf("CircleCI:\t%s%s\t%s\n", cciResp.Status.Indicator, cciTabs, cciResp.Status.Description)
		fmt.Printf("GitHub:\t\t%s%s\t%s\n", ghResp.Status.Indicator, ghTabs, ghResp.Status.Description)
	},
}

func init() {

	rootCmd.AddCommand(statusCmd)
}
