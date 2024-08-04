SHELL := bash
.ONESHELL:
.SHELLFLAGS := -euo pipefail -c
MAKEFLAGS += --warn-undefined-variables

GOLANGCI-LINT_VERSION=$(shell cat .golangci-version)
GOLANGCI-LINT_INSTALL_DIR=$(shell go env GOPATH)/bin

APP_NAME=golangci-lint-updater
WORKSPACE_FOLDER=myproject

PHONY: setup
setup: install-hooks install-lint

.PHONY: install-hooks
install-hooks:
	git config --local core.hooksPath .githooks/

.PHONY: install-lint
install-lint:
	VERSION=$(GOLANGCI-LINT_VERSION) INSTALL_DIR=$(GOLANGCI-LINT_INSTALL_DIR) scripts/install-lint.sh

.PHONY: ci
ci: shellcheck lint

.PHONY: shellcheck
shellcheck:
	scripts/shellcheck.sh

.PHONY: lint
lint:
	$(GOLANGCI-LINT_INSTALL_DIR)/golangci-lint run

.PHONY: lint-fix
lint-fix:
	$(GOLANGCI-LINT_INSTALL_DIR)/golangci-lint run --fix

.PHONY: build
build:
	APP_NAME=$(APP_NAME) scripts/build.sh

.PHONY: run
run: build rerun

.PHONY: rerun
rerun:
	APP_NAME=$(APP_NAME) bin/$(APP_NAME) $(WORKSPACE_FOLDER)

.PHONY: clean
clean:
	rm -rf bin/

.PHONY: copy-path
copy-path:
	@echo -n "$(shell pwd)/bin/$(APP_NAME)" | pbcopy
