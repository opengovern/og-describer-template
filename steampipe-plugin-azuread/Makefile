.PHONY: install build clean

build:
	GOPRIVATE="github.com/opengovern" CC=/usr/bin/musl-gcc GOPRIVATE="github.com/opengovern" GOOS=linux GOARCH=amd64 go build -v -ldflags "-linkmode external -extldflags '-static' -s -w" -o ./build/steampipe-plugin-azuread.plugin *.go

install:
	go build -o ~/.steampipe/plugins/hub.steampipe.io/plugins/turbot/azuread@latest/steampipe-plugin-azuread.plugin *.go
