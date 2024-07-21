#!/bin/bash
set -euo pipefail

find . -path ./node_modules -prune -o -iname "*.sh" -exec shellcheck {} +
