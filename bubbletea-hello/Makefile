.PHONY: default
default: all
	@:

## Artifact

ifeq ($(OS),Windows_NT)
	executable ?= teahello.exe
endif

executable ?= teahello
sources := $(wildcard *.go **/*.go)
source_main := main.go

.PHONY: all
all: $(executable)
	@:

.PHONY: clean
clean:
	$(RM) $(executable)

$(executable): $(sources)
	go build -o $(executable) $(source_main)

## Configuration

.PHONY: print-config
print-config:
	$(info OS: $(OS))
	$(info - executable: $(executable))
	$(info - sources: $(sources))
	@:

## Development

.PHONY: format
format:
	gofumpt -l -w .

.PHONY: install-tools
install-tools:
	go install mvdan.cc/gofumpt@latest

.PHONY: run
run:
	go run .

.PHONY: test
test:
	@:
	# ginkgo run -r

# .PHONY: test-watch
# test-watch:
# 	ginkgo watch -r

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: update
update:
	go get -t -u -v
