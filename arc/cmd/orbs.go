package cmd

import (
	"fmt"
	"io/ioutil"

	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

var orbsCmd = &cobra.Command{
	Use:   "orbs",
	Short: "Display what orbs your config is using.",
	Run: func(cmd *cobra.Command, args []string) {

		cciConfig, err := ioutil.ReadFile(".circleci/config.yml")
		if err != nil {
			fmt.Printf("yamlFile.Get err   #%v ", err)
		}

		yMap := make(map[interface{}]interface{})

		err = yaml.Unmarshal(cciConfig, &yMap)
		if err != nil {
			fmt.Printf(err.Error())
		}

		count := len(yMap["orbs"].(map[string]interface{}))
		fmt.Printf("This config is using %d orbs.", count)
		fmt.Printf(" Orbs is %s", yMap["orbs"])
	},
}

func init() {

	rootCmd.AddCommand(orbsCmd)
}
