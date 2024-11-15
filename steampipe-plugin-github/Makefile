STEAMPIPE_INSTALL_DIR ?= ~/.steampipe
BUILD_TAGS = netgo

.PHONY: build install clean

build: clean
	GOPRIVATE="github.com/opengovern" CC=/usr/bin/musl-gcc GOPRIVATE="github.com/opengovern" GOOS=linux GOARCH=amd64 go build -a -v -ldflags "-linkmode external -extldflags '-static' -s -w" -o ./build/steampipe-plugin-github.plugin *.go

install: build
	go build -a -o $(STEAMPIPE_INSTALL_DIR)/plugins/hub.steampipe.io/plugins/turbot/github@latest/steampipe-plugin-github.plugin -tags "${BUILD_TAGS}" *.go

clean:
	rm -rf ./build
