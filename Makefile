all: setup fmt lint test test-e2e build build-docker

.PHONY: setup
setup:  #
	sh scripts/setup.sh

.PHONY: fmt
fmt: ; @ ## Code formatter
	sh scripts/format.sh

.PHONY: lint
lint: ; @ ## Code analysis
	sh scripts/lint.sh

.PHONY: test
test:  test-unit

.PHONY: test-%
test-%: ; @ ## Run tests
	TEST_MODE=$* sh scripts/test.sh

.PHONY: build
build: build-dev

.PHONY: build-%
build-%:
	BUILD_MODE=$* sh scripts/build.sh

.PHONY: run
run: run-code;

.PHONY: run-%
run-%: ;
	RUN_MODE=$* sh scripts/run.sh

