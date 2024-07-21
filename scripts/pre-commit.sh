#!/bin/bash
set -euo pipefail

git diff --staged --name-only --exit-code -- '*.sh' || make shellcheck
git diff --staged --name-only --exit-code -- '*.go' || make lint
git diff --staged --name-only --exit-code -- '*.ts' || bun run typecheck

git diff --name-only -- '*.sh' '*.go' '*.ts' || { echo "You have unstaged changes; commit contents might not pass 'make ci'."; exit 1; }
