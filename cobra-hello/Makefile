.PHONY: default
default: all
	@:

## Environment

ifeq ($(OS),Windows_NT)
	executable ?= cobrahello.exe
endif

executable ?= cobrahello

.PHONY: debug
debug:
	$(info OS: $(OS))
	$(info - executable: $(executable))
	$(info - sources: $(sources))
	@:

## Sources

sources := $(wildcard *.go)

## Standard targets

.PHONY: all
all: $(executable)
	@:

.PHONY: clean
clean:
	$(RM) $(executable)

.PHONY: test
test:
	ginkgo

## Module targets

.PHONY: run
run:
	go run .

.PHONY: tidy
tidy:
	go mod tidy

$(executable): $(sources)
	go build $?
