#!/usr/bin/env bash

set -e

rm -rf ./completions
mkdir ./completions

for sh in bash fish zsh; do
	go run ./arc/main.go completion "$sh" >"completions/arc.$sh"
done
