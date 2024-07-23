#!/bin/sh
set -eu

find . -path ./node_modules -prune -o -iname "*.sh" -exec shellcheck {} +
