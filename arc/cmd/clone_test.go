package cmd

import (
	"bytes"
	"os"
	"strings"
	"testing"
)

func TestCloneBadInput(t *testing.T) {

	expected := "Error:  is not a valid hostname."

	output := new(bytes.Buffer)
	rootCmd.SetOut(output)
	rootCmd.SetErr(output)

	rootCmd.SetArgs([]string{"clone", "chicken"})

	rootCmd.Execute()

	if !strings.HasPrefix(output.String(), expected) {
		t.Errorf("`arc clone` with bad input should start with `%s`, result: `%s`", expected, output)
	}
}

func TestClonePass(t *testing.T) {

	// This test will modify the local filesystem. Thus, only runs in an CI
	// environment by default.
	if os.Getenv("CI") != "true" {
		t.Skip("skipping. Set envar CI=true to run.")
	}

	output := new(bytes.Buffer)
	rootCmd.SetOut(output)
	rootCmd.SetErr(output)

	rootCmd.SetArgs([]string{"clone", "https://github.com/hubci/arc.git"})

	rootCmd.Execute()

	// check to see if repo was cloned
	if _, err := os.Stat("/home/circleci/Repos/hubci/arc/.git"); os.IsNotExist(err) {
		t.Error("`arc clone` failed to clone the arc repository.")
	}
}
