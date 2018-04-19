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

This time I implemented straight through Minimax, making the big if/else branch.

I didn't have to write `opponentOf(Player) Player` this time.  I realized this time
that - based upon which branch I'm in of the if condition - I know who the opponent is.

Writing Minimax seemed pretty straightforward this time.  Negamax was introduced as a refactoring
of an already working algorithm.

What's different from prior attempts this week?

- Test-driving Minimax instead of Negamax.  I was trying to drive out Negamax one negative
  sign at a time, and that was getting confusing.
- It doesn't seem like any test cases or steps were missing this time.  Nothing makes me
  feel like I should come up with anything more than 1-level or 2-level game trees.
