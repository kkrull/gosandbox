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

We made a lot of progress with just these games:

- endGame that's already over
- openGame with pre-defined scores for each move
- twoStepGame came at the very end of our session, and was used to push recursion


Different and helpful, this time around:

- Smaller, early-stage tests, including "does not return an error" tests and "picks any non-zero" space.
- Burying all the outcome-determining logic inside a single game test double
- Having a game board where other pieces were created (or became available?) was less scary than it seemed
  for logic in a test double, and it was helpful in simplifying the test-driving


Observations:

- Test double code still feels abnormally large
- We had recursion at some point in the middle, then we lost it.  Recursion may be the only, significant piece missing from this.


Thoughts:

- Having rudimentary games does seem to be helping, so far.
- Review the minimax algorithm and identify the component behaviors / parts,
  while avoiding tying my instincts to the particular way it's expressed
- Test drive smaller parts of the algorithm and/or the games themselves?
  There was some discussion suggesting that perhaps test-driving smaller parts from the bottom up may be helpful.
- Would it help to think of the rudimentary games more formally in terms of state transitions?


## Game logic

### EndGame

There are no transitions to other games.  Any attempted move is an error.


### Quick Draw McKrull

aka King of the Mountain (without pushing)<br/>
aka First in Wins<br/>
aka Duel<br/>
aka Git on a team of thugs<br/>

[Unclaimed]

- Player claims a useless space -> Unclaimed
- Player claims the victory piece -> Player Wins


[Player Wins]

No state transitions.


### Dueling Dining Philosophers

You have to claim two tokens to win.  Empty spaces don't provide any value.

- Spaces: 2 tokens, 1 empty
- Win: Claim both tokens

The maximizing player will have to pick one of the tokens<br/>
Then the minimizing player will have to pick the other token to cause a draw<br/>
Then the game will end in a draw because both tokens have been taken and neither can get the other<br/>
