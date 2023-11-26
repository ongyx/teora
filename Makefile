## Teora Makefile

## 
## These variables can be set to customise building Teora:
## * BINARY: The path to the built binary.
## * TAGS: The tags to use to build Teora:
##     * debug.pprof: Start a live pprof server for profiling.
## 

BINARY := build/teora
TAGS :=

GOFLAGS := -tags='$(TAGS)'

help:    ## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

go:
	go build ./internal/main

native: go
native:  ## Build Teora as a native binary.

windows: export GOOS := windows
windows: export GOARCH := amd64
windows: export GOFLAGS := $(GOFLAGS) -ldflags=-H=windowsgui -o=$(BINARY).exe
windows: export CGO_ENABLED := 1
windows: export CC := x86_64-w64-mingw32-gcc
windows: go
windows: ## Build Teora on Windows.

assets:  ## Build Teora's assets.
	$(MAKE) -C internal/assets/fonts/teoran
