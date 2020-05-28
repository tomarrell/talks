.PHONY: build
build:
	pandoc -t beamer -s ./slides.md -o out.pdf

.PHONY: run
run: build
	open out.pdf

.PHONY: build-notes
build-notes:
	pandoc README.md -o readme.pdf
