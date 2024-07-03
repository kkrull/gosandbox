.PHONY: default
default: all
	@:

.PHONY: all
all:
	@:

.PHONY: clean
clean:
	@:

.PHONY: format
format:
	gofumpt -l -w .

.PHONY: install-tools
install-tools:
	go install mvdan.cc/gofumpt@latest

.PHONY: test
test:
	go test -v greeter_test.go

.PHONY: tidy
tidy:
	go mod tidy

.PHONY: update
update:
	go get -t -u -v
