all: test vet
.PHONY: all

download: go.sum
	go mod download
.PHONY: download

mod-tidy:
	go mod tidy
.PHONY: mod-tidy

generate: download
	go generate -x ./pkg/...
ifeq (, $(findstring test,$(MAKECMDGOALS))$(findstring vet,$(MAKECMDGOALS)))
	@# When not running tests or vetting, we don't want to generate mocks (vektra/mockery) as it's quite slow
else
	go generate -x ./build
endif
	gofmt -s -w .
.PHONY: generate

# Linting
vet: download generate
	gofmt -d -e -s .
	go vet ./...
	go run honnef.co/go/tools/cmd/staticcheck ./...
.PHONY: vet

vet-ci: download
	gofmt -d -e -s .
	go vet -stdmethods=false -unsafeptr=false ./...
	go run honnef.co/go/tools/cmd/staticcheck ./...
.PHONY: vet-ci

# Testing
test: download generate
	go test -v ./...
.PHONY: test

TEST_OUT_DIR ?= .
COV_OUT_DIR ?= $(TEST_OUT_DIR)

test-ci: download
	go run gotest.tools/gotestsum --junitfile $(TEST_OUT_DIR)/report.xml -- -coverprofile=$(TEST_OUT_DIR)/coverage.out -v ./...
.PHONY: test-ci

COV_IN_FILE ?= $(TEST_OUT_DIR)/coverage.out

coverage-ci: download
	go tool cover -html=$(COV_IN_FILE) -o $(COV_OUT_DIR)/coverage.html
	go run github.com/t-yuki/gocover-cobertura < $(COV_IN_FILE) > $(COV_OUT_DIR)/cobertura.xml
.PHONY: coverage-ci

# Build
build: download generate
	go build -x -v -work ./pkg/entry
.PHONY: build

# Clean
clean:
	find . -name '*.gen.*' -delete
	rm -f ./pkg/gdraw/*.json ./pkg/gdraw/*.h
.PHONY: clean
