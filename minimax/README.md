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
- Still struggling to drive out more rules.  I'm currently working on not just ending a game, but on determining who the
  winner is.  I think I'll have to switch my notion of capture the flag to having two flags for two opposing bases
