SHELL := bash
.ONESHELL:
.SHELLFLAGS := -eu -o pipefail -c
MAKEFLAGS += --warn-undefined-variables

GOLANGCI_LINT_VERSION=$(shell cat .golangci.version)
GOLANGCI_LINT_INSTALL_DIR=$(shell go env GOPATH)/bin

.PHONY: lint-install
lint-install:
	VERSION=$(GOLANGCI_LINT_VERSION) INSTALL_DIR=$(GOLANGCI_LINT_INSTALL_DIR) scripts/install-lint.sh

.PHONY: lint
lint:
	$(GOLANGCI_LINT_INSTALL_DIR)/golangci-lint run

.PHONY: lint-fix
lint-fix:
	$(GOLANGCI_LINT_INSTALL_DIR)/golangci-lint run --fix

.PHONY: shellcheck
shellcheck:
	scripts/shellcheck.sh