.EXPORT_ALL_VARIABLES:

SHELL=/bin/bash -o pipefail

build:
	@go build ./index_buildpack.go