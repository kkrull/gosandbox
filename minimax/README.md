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

- Moving to this project and avoiding the `GOPATH` issues was essential.
- `Board` should probably just be `Game`.  Thinking of games in the abstract and passing in functions that contain the
  game's rules helps me to focus on the Minimax algorithm
- The Minimax logic is still incomplete (missing recursion)
- The Minimax logic is also heading in an odd direction.  It needs to select among moves instead of picking the first
  terminal state it finds.
