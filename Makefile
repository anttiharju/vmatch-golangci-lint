.PHONY: install-linter
install-linter:
	scripts/install-lint.sh

.PHONY: shellcheck
shellcheck:
	scripts/shellcheck.sh
