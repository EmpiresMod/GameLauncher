#!/usr/bin/make -f

SHELL=/bin/sh
bin=bin
name=launcher

all: build

build:
	go build -v -o bin/$(name)

windows:
	CC=i686-w64-mingw32-gcc-4.6 GOOS=windows GOARCH=386 CGO_ENABLED=1 go build -v -ldflags="-H=windowsgui" -o="bin/$(name).exe"

windows-debug:
	CC=i686-w64-mingw32-gcc-4.6 GOOS=windows GOARCH=386 CGO_ENABLED=1 go build -v -o="bin/$(name).exe"

clean:
	go clean -x

remove:
	go clean -i
