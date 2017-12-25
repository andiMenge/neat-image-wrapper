#!/bin/bash
name="neatbatch"
set -ex

printf "\n## BUILD RELEASES ##\n"

echo " build macOS binary..."
GOOS=darwin go build -o $name ./main.go

echo "build windows binary..."
GOOS=windows go build -o $name.exe ./main.go
