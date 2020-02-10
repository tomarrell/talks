.PHONY: build
build:
	pandoc -t beamer -s ./slides.md -o out.pdf

.PHONY: run
run: build
	open out.pdf
