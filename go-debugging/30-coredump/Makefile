up:
	docker run -p 8080:8080 -v /Users/tom/Documents/talks/go-debugging/30-coredump:/coredump -it golang:buster

ulimit: ## Run this manually, with Make it will not work
	ulimit -c unlimited

build:
	go build .

dump: ulimit build
	GOTRACEBACK=crash ./coredump

dlv:
	go get -u github.com/go-delve/delve/cmd/dlv

inspect:
	dlv core ./coredump ./core
