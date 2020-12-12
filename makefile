.EXPORT_ALL_VARIABLES:

SHELL=/bin/bash -o pipefail

build:
	@go build -o bin/ ./...

test:
	@go test ./...

resetdb:
	@dropdb cnb-registry-api-dev
	@createdb cnb-registry-api-dev
