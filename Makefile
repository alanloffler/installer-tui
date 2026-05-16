.PHONY: run
run:
	@go run .
run-no-auth:
	@GH_CONFIG_DIR=$(shell mktemp -d) go run .
run-no-gh:
	@go build -o /tmp/installer-tui-test . && PATH=/usr/bin:/bin /tmp/installer-tui-test
