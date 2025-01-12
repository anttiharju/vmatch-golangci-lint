#!/usr/bin/env dash
set -eu

SHA=$(git rev-parse HEAD)
if [ -n "$(git status --porcelain)" ]; then
    BUILDID="$SHA-dirty"
else
    BUILDID="$SHA"
fi

go build -C cmd/$APP_NAME -ldflags "-s -w -buildid=$BUILDID -X github.com/anttiharju/vmatch/pkg/exit.appName=$APP_NAME" -o "$(pwd)/$APP_NAME"
