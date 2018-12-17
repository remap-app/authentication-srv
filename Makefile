NAME := authentication
VERSION := $(shell git describe --tags --abbrev=0)
REVISION := $(shell git rev-parse --short HEAD)
LDFLAGS := -X 'main.version=$(VERSION)' \
	-X 'main.revision=$(REVISION)'

## Setup

setup:
	go get github.com/golang/lint/golint


