# Hello World in Go

Test-driving a simple Go function with `go1.10 darwin/amd64`. 


## Source and dependencies

The code must be in the location defined in your [`GOPATH`](https://golang.org/doc/code.html#GOPATH). 
Use `go get -t` to fetch source code for dependencies.


## Testing

Tests are run with `testing`, organized with [Ginkgo](https://github.com/onsi/ginkgo), and assert themselves with
[Gomega](http://onsi.github.io/gomega/).

- Test once with `ginkgo` or `go test`.  This shows dots for each test kind of like RSpec's default reporter.
- Show test names as they run with `ginkgo -v` or `go test -ginkgo.v`.
- Run tests automatically when files are modified with `ginkgo watch`.

