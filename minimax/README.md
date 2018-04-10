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

- Negamax is a lot faster
- I'm getting mixed up on the sign flipping
- The last test felt bit too large of a step, because I had to flip the terminal scores and the result of the Negamax

