# Minimax Kata in Go

## Setup

This is being developed on Go version `go1.10 darwin/amd64`.

```bash
$ cd $(go env GOPATH)/github.com/kkrull/gosandbox/minimax
$ go get -t
```

## Testing

```bash
$ ginkgo watch
```

## How this was made

```bash
$ ginkgo bootstrap
$ ginkgo generate minimax
```

## Thoughts

- Formalizing the notion of having different Game types is helping to write the tests
- It's easy to introduce value types by using placeholder interfaces, which prevents refactoring later when types change
  from primitives to types with meaning
- Still struggling to drive out more rules.  It's hard not to get caught up on the semantics of a simple game that has
  to work in exactly 1 scenario.

