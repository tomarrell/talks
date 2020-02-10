.PHONY: up
up:
	docker run -v /Users/tom/Documents/talks/go-debugging/3-coredump:/coredump -it golang:buster

.PHONY: ulimit
ulimit:
	ulimit -c unlimited

.PHONY: build
build:
	go build .

.PHONY: dump
dump: ulimit build
	GOTRACEBACK=crash ./coredump

.PHONY: delve
delve:
	go get -u github.com/go-delve/delve/cmd/dlv

.PHONY: inspect
inspect:
	dlv core ./coredump ./core
