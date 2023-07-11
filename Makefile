BINARY_NAME := route53-to-zonefile
VERSION := $(shell git describe --tags --always 2>/dev/null || echo "v0.0.0")
PLATFORMS := linux darwin windows
ARCHITECTURES := amd64 386
BUILD_OUTPUT := dist

.PHONY: all clean build release

all: clean build

clean:
	rm -rf $(BUILD_OUTPUT)

build:
	go build -o $(BINARY_NAME) -ldflags "-X main.version=$(VERSION)"

release: clean
	mkdir -p $(BUILD_OUTPUT)
	$(foreach os, $(PLATFORMS), \
		$(foreach arch, $(ARCHITECTURES), \
			GOOS=$(os) GOARCH=$(arch) go build -o $(BUILD_OUTPUT)/$(BINARY_NAME)-$(os)-$(arch)$(if $(findstring windows,$(os)),.exe) -ldflags "-X main.version=$(VERSION)"; \
		) \
	)
