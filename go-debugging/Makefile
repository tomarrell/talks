.PHONY: build
build:
	pandoc -t beamer -s ./slides.md -o slides.pdf

.PHONY: run
run: build
	open slides.pdf

.PHONY: build-notes
build-notes:
	pandoc README.md -o notes.pdf
