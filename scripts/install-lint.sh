#!/usr/bin/env bash

version=$(cat .golangci.version)
install_dir=$(go env GOPATH)/bin

mkdir -p "$install_dir"
cd "$install_dir" || exit
curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b . "$version"
