#!/bin/bash

find . -iname "*.sh" -exec shellcheck {} +
