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

This time I wrote vanilla Minimax and refactored my way into Negamax in order to collapse
the if/else branch in Minimax.

I think it largely works as a refactoring, but it does require some insight that was not
initially intuitive to me so that you can figure out where to keep flipping negatives.

At the point that I introduced the refactoring, I hadn't switched players.  The refactoring
introduced the polarity flip.  I suppose I missed a test in Megamax in taking 3 turns
(recurring twice) instead of 2.  