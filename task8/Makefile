SRC := main.go
MODULE := github.com/alex0ps/go-yandex/task8

.PHONY: all
all: build

go.mod:
	go mod init $(MODULE)

.PHONY: build
build: go.mod
	go build ./
