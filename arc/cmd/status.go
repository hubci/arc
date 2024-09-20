package cmd

import (
	"fmt"

	"github.com/hubci/arc/arc/statuses"
	"github.com/spf13/cobra"
)

var (
	cciFl bool

	statusCmd = &cobra.Command{
		Use:   "status <name>",
		Short: "Provides the status page result for the provided name",
		Args:  cobra.ExactArgs(1),
		Run: func(cmd *cobra.Command, args []string) {

			statusPage, err := statuses.Page(args[0])
			if err != nil {
				fmt.Println(err)
				return
			}

			client := New()

			err = statusPage.Fetch(client.c)
			if err != nil {
				fmt.Println(err)
				return
			}

			status, err := statusPage.Status()
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(status)
		},
	}
)

func init() {

	// Temporary flag to hide non-CircleCI related statuses
	statusCmd.Flags().BoolVar(&cciFl, "cci", false, "only show CircleCI related statuses")
	statusCmd.Flags().MarkHidden("cci")
	rootCmd.AddCommand(statusCmd)
}
