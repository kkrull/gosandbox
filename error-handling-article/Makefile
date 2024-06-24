sources := \
	grpfn/grpfn.go \
	mapfn/mapfn.go \
	maphof/maphof.go \
	parseifs/parse_ifs.go \
	tcplisten/tcp_listen.go \
	wrapfn/wrapfn.go

.PHONY: default
default: run
	@:

.PHONY: all
all:
	@:

.PHONY: clean
clean:
	@:

.PHONY: run
run:
	$(foreach source,$(sources),go run $(source);)

.PHONY: test
test:
	@:

.PHONY: tidy
tidy:
	go mod tidy
