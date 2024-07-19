#!/usr/bin/env bash

find . -iname "*.sh" | xargs shellcheck
