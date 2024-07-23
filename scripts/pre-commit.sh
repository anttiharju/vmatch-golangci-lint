#!/bin/sh
set -eu

changes_to() {
    for filetype in "$@"; do
        if git diff --name-only -- "$filetype" || git diff --cached --name-only -- "$filetype"; then
            return 0
        fi
    done
    return 1
}

if changes_to "*.sh"; then
    make shellcheck
fi

if changes_to "*.go"; then
    make lint
fi

if changes_to "*.ts"; then
    bun run typecheck
fi

git diff --name-only --exit-code -- '*.sh' '*.go' '*.ts' || {
	echo "You have unstaged changes; commit contents might not pass 'make ci'."
	exit 1
}
