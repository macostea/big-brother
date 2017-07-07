DATE    ?= $(shell date +%FT%T%z)
VERSION ?= $(shell git describe --tags --always --dirty --match=v* 2> /dev/null || \
			cat $(CURDIR)/.version 2> /dev/null || echo v0)

.PHONY: all agent master clean get-deps
all: clean get-deps agent master

agent:
		go build \
		-tags release \
		-ldflags '-X agent/cmd.Version=$(VERSION) -X agent/cmd.BuildDate=$(DATE)' \
		-o out/bin/agent agent/main/main.go

master:
		go build \
		-tags release \
		-ldflags '-X master/cmd.Version=$(VERSION) -X master/cmd.BuildDate=$(DATE)' \
		-o out/bin/master master/main/main.go

clean:
		rm -rf out/
		echo Clean complete

get-deps:
		go get -t ./...
