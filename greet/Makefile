default: build

## Environment

ifeq ($(OS),Windows_NT)
	executable ?= greet.exe
endif

executable ?= greet

.PHONY: debug
debug:
	$(info OS: $(OS))
	$(info - executable: $(executable))
	$(info - sources: $(sources))
	@:

## Sources

sources := $(wildcard *.go)

## Standard targets

.PHONY: clean
clean:
	$(RM) $(executable)

## Greeter

.PHONY: build
build: $(executable)
	@:

$(executable): $(sources)
	go build $?

.PHONY: run
run:
	go run .

.PHONY: tidy
tidy:
	go mod tidy
