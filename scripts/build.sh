#!/bin/sh
set -eu

ROOT=$(pwd)
mkdir -p "$ROOT/bin"
go build -C src -ldflags "-s -w -buildid=" -o "$ROOT/bin/$APP_NAME"
