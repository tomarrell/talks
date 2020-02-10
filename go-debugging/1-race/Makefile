.PHONY: build
build:
	go build race.go

.PHONY: run
run: build
	for i in {1..1000}; do ./race; done
