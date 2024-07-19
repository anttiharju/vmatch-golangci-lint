#!/usr/bin/env bash

find . -iname "*.sh" -print0 | xargs shellcheck
