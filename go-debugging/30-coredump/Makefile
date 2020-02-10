up:
	docker run -v /Users/tom/Documents/talks/go-debugging/3-coredump:/coredump -it golang:buster

ulimit:
	ulimit -c unlimited

build:
	go build .

dump: ulimit build
	GOTRACEBACK=crash ./coredump

delve:
	go get -u github.com/go-delve/delve/cmd/dlv

inspect:
	dlv core ./coredump ./core
