## Teora Makefile

## 
## These variables can be set to customise building Teora:
## * BINARY: The path to the built binary.
## * TAGS: The tags to use to build Teora:
##     * debug: Enable debug mode.
##     * pprof: Start a live pprof server for profiling.
## * FLAGS: Flags to pass to the 'go build' command.
## 

BINARY := build/teora
TAGS := debug
FLAGS :=

.PHONY: bootstrap clean

help:       ## Show this help.
	@sed -ne '/@sed/!s/## //p' $(MAKEFILE_LIST)

native:     ## Build Teora as a native binary.
	go build -tags=$(TAGS) -o $(BINARY) ./internal/main

windows: export GOOS := windows
windows: export GOARCH := amd64
windows:    ## Build Teora as a Windows console app.
	go build -tags=$(TAGS) $(FLAGS) -o $(BINARY).exe ./internal/main

windowsgui: export FLAGS := $(FLAGS) -ldflags -H=windowsgui
windowsgui: windows
windowsgui: ## Build Teora as a Windows GUI app.

assets:     ## Build Teora's assets.
	$(MAKE) -C internal/assets/fonts/teoran
