# Ginkgo v2 Sandbox

Try Ginkgo since it has been upgraded to use Go Modules.

## Install the Ginkgo CLI

- Install go and make sure `$(go env GOPATH)/bin` is on your `PATH`.
- Initialize a module
- Install the ginkgo CLI

```sh
go get github.com/onsi/ginkgo/v2/ginkgo
go install github.com/onsi/ginkgo/v2/ginkgo
```

## Bootstrap Ginkgo entry point

```sh
ginkgo bootstrap
go get -t # Might be unnecessary?
go mod tidy
```

The test suite now runs, with 0 specs:

```sh
ginkgo
go test -v ./... # Alternative
```

## Generate test file

```sh
ginkgo generate greeter
```

Add a spec to the generated file, and now it runs.
