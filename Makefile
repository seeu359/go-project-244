COVERAGE_MIN = 90
PACKAGES = $(shell go list ./... | grep -v /cmd/ | paste -sd,)

build:
	go build -o bin/gendiff ./cmd/gendiff

lint:
	golangci-lint run

test:
	go test ./...

test-coverage:
	go test -covermode=count -coverpkg=$(PACKAGES) -coverprofile=coverage.out ./...
	go tool cover -func=coverage.out
	@go tool cover -func=coverage.out | awk -v min=$(COVERAGE_MIN) '/^total:/ { gsub(/%/, "", $$3); if ($$3 + 0 < min) { print "coverage " $$3 "% is below minimum " min "%"; exit 1 } }'

.PHONY: build lint test test-coverage
