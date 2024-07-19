#!/usr/bin/env bash

which shellcheck
find . -iname "*.sh" -exec shellcheck {} +
