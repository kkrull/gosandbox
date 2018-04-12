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

- 5 test cases.  _Am I done already???_
- Starting with the Minimax function and introducing Negamax later was helpful
  in terms of me keeping track of the expected scores.
- Switching from Minimax to introduce Negamax happened during a refactoring step.
  The refactoring wasn't exactly addressing a cleanliness issue with the existing algorithm,
  it was really preparing to make the next test easier.  Is that cheating?
- Naming the tests more carefully was helpful, although I still had some confusion in 1 case
  about the expected score for the root node.
