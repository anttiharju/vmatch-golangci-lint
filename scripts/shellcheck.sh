#!/bin/bash
set -euco pipefail

# change
find . -iname "*.sh" -exec shellcheck {} +
