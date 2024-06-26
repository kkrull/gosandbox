.PHONY: default
default: all
	@:

## Artifact

ifeq ($(OS),Windows_NT)
	executable ?= cobrahello.exe
endif

executable ?= cobrahello
sources := $(wildcard *.go)

.PHONY: all
all: $(executable)
	@:

.PHONY: clean
clean:
	$(RM) $(executable)

$(executable): $(sources)
	go build -o $(executable) $?

## Configuration

.PHONY: print-config
print-config:
	$(info OS: $(OS))
	$(info - executable: $(executable))
	$(info - sources: $(sources))
	@:

## Development

.PHONY: install-tools
install-tools:
	go install github.com/spf13/cobra-cli@latest

.PHONY: run
run:
	go run .

.PHONY: test
test:
	ginkgo

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: update
update:
	go get -t -u -v