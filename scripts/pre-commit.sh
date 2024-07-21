#!/bin/bash
set -euo pipefail

make ci
git diff --quiet || { echo "You have unstaged changes; commit contents might not pass 'make ci'."; exit 1; }
