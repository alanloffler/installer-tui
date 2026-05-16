.PHONY: run
run:
	go run .
run-no-auth:
	GH_CONFIG_DIR=$(shell mktemp -d) go run .
run-no-gh:
	PATH=/usr/bin:/bin go run .
