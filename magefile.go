//go:build mage
// +build mage

package main

import (
	"github.com/magefile/mage/sh"
)

// Install Arc by building a binary from the current source and placing it in
// the Go PATH
func Install() error {

	if err := sh.Run("go", "mod", "download"); err != nil {
		return err
	}

	return sh.Run("go", "install", "./arc")
}

// Remove the Arc binary created from the Install command
func Remove() error {

	return sh.Run("go", "clean", "-i", "github.com/hubci/arc/arc")
}

func Test() error {

	return sh.Run("gotestsum", "./...")
}

func TestCI() error {

	return sh.Run("gotestsum", "--junitfile=junit/unit-tests.xml", "--", "-coverprofile=coverage.txt", "-covermode=atomic", "./...")
}
