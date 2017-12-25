#!/bin/bash
name="neatbatch"

printf "\n## BUILD RELEASES ##"

echo " build macOS binary..."
GOOS=darwin go build -o $name main.go

echo "build windows binary..."
GOOS=windows go build -o $name main.go
