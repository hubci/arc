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

type sStatus struct {
	Updated    time.Time `json:"updated"`
	Status     string    `json:"status"`
	StatusCode int       `json:"status_code"`
}

type sResult struct {
	StatusOverall *sStatus `json:"status_overall"`
}

type sResponse struct {
	Result *sResult `json:"result"`
}

var (
	cciFl bool

	statusCmd = &cobra.Command{
		Use:   "status",
		Short: "Provides the status page results for various DevOps services",
		Run: func(cmd *cobra.Command, args []string) {

			var cciResp *spResponse
			var cfResp *spResponse
			var ghResp *spResponse
			var gitlabResp *sResponse
			var linodeResp *spResponse
			var doResp *spResponse

			cciURL := "https://status.circleci.com/api/v2/status.json"
			cfURL := "https://www.cloudflarestatus.com/api/v2/status.json"
			ghURL := "https://www.githubstatus.com/api/v2/status.json"
			gitlabURL := "https://status.gitlab.com/1.0/status/5b36dc6502d06804c08349f7"
			linodeURL := "https://status.linode.com/api/v2/status.json"
			doURL := "https://status.digitalocean.com/api/v2/status.json"

			client := New()

			errCCI := client.getJSON(cciURL, &cciResp)
			errCF := client.getJSON(cfURL, &cfResp)
			errGH := client.getJSON(ghURL, &ghResp)
			errGitlab := client.getJSON(gitlabURL, &gitlabResp)
			errLinode := client.getJSON(linodeURL, &linodeResp)
			errDO := client.getJSON(doURL, &doResp)

			if errCCI != nil {
				cciResp.Status.Indicator = "can't connect"
			}

			if errCF != nil {
				cfResp.Status.Indicator = "can't connect"
			}

			if errGH != nil {
				ghResp.Status.Indicator = "can't connect"
			}

			if errGitlab != nil {
				gitlabResp.Result.StatusOverall.Status = "can't connect"
			}

			if errLinode != nil {
				linodeResp.Status.Indicator = "can't connect"
			}

			if errDO != nil {
				doResp.Status.Indicator = "can't connect"
			}

			var cciTabs = ""
			if cciResp.Status.Indicator == "none" {
				cciResp.Status.Indicator = ""
				cciTabs = "\t"
			}

			var cfTabs = ""
			if cfResp.Status.Indicator == "none" {
				cfResp.Status.Indicator = ""
				cfTabs = "\t"
			}

			var ghTabs = ""
			if ghResp.Status.Indicator == "none" {
				ghResp.Status.Indicator = ""
				ghTabs = "\t"
			}

			var linodeTabs = ""
			if linodeResp.Status.Indicator == "none" {
				linodeResp.Status.Indicator = ""
				linodeTabs = "\t"
			}

			var doTabs = ""
			if doResp.Status.Indicator == "none" {
				doResp.Status.Indicator = ""
				doTabs = "\t"
			}

			fmt.Println("Reporting status page results...")
			fmt.Println("")
			fmt.Printf("CircleCI:\t%s%s\t%s\n", cciResp.Status.Indicator, cciTabs, cciResp.Status.Description)
			fmt.Printf("GitHub:\t\t%s%s\t%s\n", ghResp.Status.Indicator, ghTabs, ghResp.Status.Description)
			fmt.Printf("Gitlab:\t\t\t\t%s\n", gitlabResp.Result.StatusOverall.Status)
			if !cciFl {
				fmt.Printf("Cloudflare:\t\t%s%s\t%s\n", cfResp.Status.Indicator, cfTabs, cfResp.Status.Description)
				fmt.Printf("Linode:\t\t%s%s\t%s\n", linodeResp.Status.Indicator, linodeTabs, linodeResp.Status.Description)
				fmt.Printf("DigitalOcean:\t%s%s\t%s\n", doResp.Status.Indicator, doTabs, doResp.Status.Description)
			}
		},
	}
)

func init() {

	// Temporary flag to hide non-CircleCI related statuses
	statusCmd.Flags().BoolVar(&cciFl, "cci", false, "only show CircleCI related statuses")
	statusCmd.Flags().MarkHidden("cci")
	rootCmd.AddCommand(statusCmd)
}
