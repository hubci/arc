package cmd

import (
	"fmt"
	"io/ioutil"
	"strings"

	"golang.org/x/mod/semver"

	"github.com/algolia/algoliasearch-client-go/v3/algolia/search"
	"github.com/goccy/go-yaml"
	"github.com/spf13/cobra"
)

// a partial example of a CircleCI configuration file
type cciConfig struct {
	Orbs map[string]string `yaml:"orbs"`
}

// a CircleCI orb
type orb struct {
	name          string
	namespace     string
	version       string
	LatestVersion string `json:"version,omitempty"`
	ObjectID      string `json:"objectID,omitempty"`
}

func orbParse(definition string) orb {

	slash := strings.Index(definition, "/")
	at := strings.Index(definition, "@")

	return orb{
		namespace: definition[0:slash],
		name:      definition[slash+1 : at],
		version:   definition[at+1:],
	}
}

var orbsCmd = &cobra.Command{
	Use:   "orbs",
	Short: "Display what orbs your config is using.",
	Run: func(cmd *cobra.Command, args []string) {

		configFile, err := ioutil.ReadFile(".circleci/config.yml")
		if err != nil {
			fmt.Printf("yamlFile.Get err   #%v ", err)
		}

		var config cciConfig

		err = yaml.Unmarshal(configFile, &config)
		if err != nil {
			fmt.Printf(err.Error())
		}

		count := len(config.Orbs)

		if count == 0 {
			fmt.Println("This config doesn't use any orbs.")
			return
		}

		fmt.Printf("This config is using %d orbs.\n\n", count)

		client := search.NewClient("U0RXNGRK45", "0efdf785bc8c215d9f5fd7fdece50795")
		index := client.InitIndex("orbs-prod")

		for _, orbStr := range config.Orbs {

			anOrb := orbParse(orbStr)
			err := index.GetObject(anOrb.namespace+"/"+anOrb.name, &anOrb)
			if err != nil {
				fmt.Printf(err.Error())
			}

			fmt.Printf("%s/%s: v%s", anOrb.namespace, anOrb.name, anOrb.version)

			if semver.Compare("v"+anOrb.version, "v"+anOrb.LatestVersion) == -1 {
				fmt.Printf(" (latest v%s)", anOrb.LatestVersion)
			}

			//intentionally empty
			fmt.Println()
		}
	},
}

func init() {

	rootCmd.AddCommand(orbsCmd)
}
