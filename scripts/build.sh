#!/bin/sh
set -eu

ROOT=$(pwd)
mkdir -p "$ROOT/bin"
go build -C cmd -o "$ROOT/bin/$APP_NAME"
chmod +x "$ROOT/bin/$APP_NAME"
