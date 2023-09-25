BASH ?= $(shell command -v bash 2> /dev/null)

# Testing
GO_TEST_FLAGS ?= -v -race -count=1
GOTESTFMT_FLAGS ?=
# More flags can be injected into the test by using the EXTRA_TEST_FLAGS environment variable

## Run unit tests
.PHONY: unittest
unittest:
	GO_TEST_FLAGS="$(GO_TEST_FLAGS)" GOTESTFMT_FLAGS="$(GOTESTFMT_FLAGS)" $(BASH) -xe ./scripts/test.sh
