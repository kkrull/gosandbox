default: run

.PHONY: run
run:
	go run grpfn/grpfn.go
	go run mapfn/mapfn.go
	go run maphof/maphof.go
	go run parseifs/parse_ifs.go
	go run tcplisten/tcp_listen.go
	go run wrapfn/wrapfn.go

.PHONY: tidy
tidy:
	go mod tidy
