PROJECT = project-layout-go
TAG = $(shell git tag -l --points-at HEAD)
AUTHOR = $(shell git log --pretty=format:"%an"|head -n 1)
VERSION = $(shell git rev-list HEAD | head -1)
BUILD_INFO = $(shell git log --pretty=format:"%s" | head -1)
BUILD_DATE = $(shell date +%Y-%m-%d\ %H:%M:%S)
CUR_PWD := $(shell pwd)

export GO111MODULE=on
export GOPROXY=https://goproxy.cn,direct
export GOSUMDB=off

LD_FLAGS='-X "$(PROJECT)/version.TAG=$(TAG)" -X "$(PROJECT)/version.VERSION=$(VERSION)" -X "$(PROJECT)/version.AUTHOR=$(AUTHOR)" -X "$(PROJECT)/version.BUILD_INFO=$(BUILD_INFO)" -X "$(PROJECT)/version.BUILD_DATE=$(BUILD_DATE)"'

default: build

TARGETLIST = $(shell ls cmd)

.PHONY: build

build: ${TARGETLIST}

${TARGETLIST}: %:
	go build -v -ldflags $(LD_FLAGS) -gcflags "-N" -o ./bin/$@ cmd/$@/main.go
