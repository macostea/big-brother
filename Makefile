DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

.PHONY: all clean get-deps test
all: get-deps test out/bin/agent out/bin/master

out/bin/agent:
		go build \
		-tags release \
		-ldflags '-X agent/cmd.Version=$(VERSION) -X agent/cmd.BuildDate=$(DATE)' \
		-o out/bin/agent/agent agent/main/main.go
		cp agent/config/config.yaml out/bin/agent

out/bin/master:
		go build \
		-tags release \
		-ldflags '-X master/cmd.Version=$(VERSION) -X master/cmd.BuildDate=$(DATE)' \
		-o out/bin/master/master master/main/main.go
		cp master/config/servers.yaml out/bin/master
clean:
		rm -rf out/
		echo Clean complete

get-deps:
		go get -t ./...

test:
		go test ./... -cover