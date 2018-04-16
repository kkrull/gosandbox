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

### Step 1 -- introduce function

commit 468a6db2d6b25911152ddb8a430ff944d72e3c5a
Author: Kyle Krull <kkrull@8thlight.com>
Date:   Mon Apr 16 09:30:47 2018 -0500

    Exists

diff --git a/minimax/minimax.go b/minimax/minimax.go
index 2757f81..3149d2e 100644
--- a/minimax/minimax.go
+++ b/minimax/minimax.go
@@ -1,2 +1,5 @@
 package minimax

+func Minimax() int {
+       return 0
+}
diff --git a/minimax/minimax_test.go b/minimax/minimax_test.go
index ff84937..2f4f86d 100644
--- a/minimax/minimax_test.go
+++ b/minimax/minimax_test.go
@@ -2,11 +2,13 @@ package minimax_test

 import (
        . "github.com/onsi/ginkgo"
-       //. "github.com/onsi/gomega"
+       . "github.com/onsi/gomega"

-       //. "github.com/kkrull/gosandbox/minimax"
+       . "github.com/kkrull/gosandbox/minimax"
 )

 var _ = Describe("Minimax", func() {
-
+       It("scores a game ending in a draw as 0", func() {
+               Expect(Minimax()).To(Equal(0))
+       })
 })


### Step 2 -- expand function to have the target interface

diff --git a/minimax/minimax.go b/minimax/minimax.go
index 3149d2e..7f9367c 100644
--- a/minimax/minimax.go
+++ b/minimax/minimax.go
@@ -1,5 +1,13 @@
 package minimax

-func Minimax() int {
+func Minimax(game Game, player Player) int {
        return 0
 }
+
+type Game interface {
+
+}
+
+type Player struct {
+       Name string
+}
\ No newline at end of file
diff --git a/minimax/minimax_test.go b/minimax/minimax_test.go
index 2f4f86d..484f87a 100644
--- a/minimax/minimax_test.go
+++ b/minimax/minimax_test.go
@@ -7,8 +7,17 @@ import (
        . "github.com/kkrull/gosandbox/minimax"
 )

+var (
+       max = Player{Name: "max"}
+)
+
 var _ = Describe("Minimax", func() {
        It("scores a game ending in a draw as 0", func() {
-               Expect(Minimax()).To(Equal(0))
+               game := FakeGame{}
+               Expect(Minimax(game, max)).To(Equal(0))
        })
 })
+
+type FakeGame struct {
+
+}


### Step 3 -- refactor to force Game#IsOver check

commit 859f694b094128dbb7a5915d32bccb4f9f846665
Author: Kyle Krull <kkrull@8thlight.com>
Date:   Mon Apr 16 09:32:52 2018 -0500

    Checking that the game is over

diff --git a/minimax/minimax.go b/minimax/minimax.go
index 7f9367c..98ce8a2 100644
--- a/minimax/minimax.go
+++ b/minimax/minimax.go
@@ -1,11 +1,15 @@
 package minimax

 func Minimax(game Game, player Player) int {
-       return 0
+       if game.IsOver() {
+               return 0
+       }
+
+       panic("no result")
 }

 type Game interface {
-
+       IsOver() bool
 }

 type Player struct {
diff --git a/minimax/minimax_test.go b/minimax/minimax_test.go
index 484f87a..5d70966 100644
--- a/minimax/minimax_test.go
+++ b/minimax/minimax_test.go
@@ -13,11 +13,16 @@ var (

 var _ = Describe("Minimax", func() {
        It("scores a game ending in a draw as 0", func() {
-               game := FakeGame{}
+               game := FakeGame{over: true}
                Expect(Minimax(game, max)).To(Equal(0))
        })
 })

 type FakeGame struct {
+       over bool
+}

+func (game FakeGame) IsOver() bool {
+       return game.over
 }
+


### Step 4 -- Score for max victory

diff --git a/minimax/minimax.go b/minimax/minimax.go
index 98ce8a2..0ef23ed 100644
--- a/minimax/minimax.go
+++ b/minimax/minimax.go
@@ -1,7 +1,9 @@
 package minimax

 func Minimax(game Game, player Player) int {
-       if game.IsOver() {
+       if game.FindWinner() == game.MaximizingPlayer() {
+               return 1
+       } else if game.IsOver() {
                return 0
        }

@@ -10,8 +12,10 @@ func Minimax(game Game, player Player) int {

 type Game interface {
        IsOver() bool
+       FindWinner() Player
+       MaximizingPlayer() Player
 }


### Step 5 -- Score for min victory

diff --git a/minimax/minimax.go b/minimax/minimax.go
index 0ef23ed..2849bf8 100644
--- a/minimax/minimax.go
+++ b/minimax/minimax.go
@@ -3,6 +3,8 @@ package minimax
 func Minimax(game Game, player Player) int {
        if game.FindWinner() == game.MaximizingPlayer() {
                return 1
+       } else if game.FindWinner() == game.MinimizingPlayer() {
+               return -1
        } else if game.IsOver() {
                return 0
        }
@@ -14,6 +16,7 @@ type Game interface {
        IsOver() bool
        FindWinner() Player
        MaximizingPlayer() Player
+       MinimizingPlayer() Player
 }


### Step 6 -- Introduce opponent search and polarity

diff --git a/minimax/minimax.go b/minimax/minimax.go
index 2849bf8..1955556 100644
--- a/minimax/minimax.go
+++ b/minimax/minimax.go
@@ -1,15 +1,20 @@
 package minimax

 func Minimax(game Game, player Player) int {
+       return Negamax(game, 1)
+}
+
+func Negamax(game Game, polarity int) int {
        if game.FindWinner() == game.MaximizingPlayer() {
-               return 1
+               return 1 * polarity
        } else if game.FindWinner() == game.MinimizingPlayer() {
-               return -1
+               return -1 * polarity
        } else if game.IsOver() {
                return 0
        }

-       panic("no result")
+       nextGame := game.NextGames()[0]
+       return -Negamax(nextGame, -1 * polarity)
 }

 type Game interface {
@@ -17,6 +22,7 @@ type Game interface {
        FindWinner() Player
        MaximizingPlayer() Player
        MinimizingPlayer() Player
+       NextGames() []Game
 }


### Step 7 -- Try to force minimizing player negation of negamax

**Work on this step to find a test that forces this to happen**

commit 19c9ac96466dcda944496ca378fa756aa0c0f1ab
Author: Kyle Krull <kkrull@8thlight.com>
Date:   Mon Apr 16 09:45:23 2018 -0500

    Passes, but didn't really force me to write this

diff --git a/minimax/minimax.go b/minimax/minimax.go
index 1955556..34dadb6 100644
--- a/minimax/minimax.go
+++ b/minimax/minimax.go
@@ -1,7 +1,11 @@
 package minimax

 func Minimax(game Game, player Player) int {
-       return Negamax(game, 1)
+       if player == game.MaximizingPlayer() {
+               return Negamax(game, 1)
+       } else {
+               return -Negamax(game, -1)
+       }
 }

 func Negamax(game Game, polarity int) int {
diff --git a/minimax/minimax_test.go b/minimax/minimax_test.go
index 012b72c..955a6b0 100644
--- a/minimax/minimax_test.go
+++ b/minimax/minimax_test.go
@@ -33,6 +33,12 @@ var _ = Describe("Minimax", func() {
                        FakeGame{isOver: true, winner: min}}}
                Expect(Minimax(game, max)).To(Equal(-1))
        })
+
+       It("scores an unfinished game that the maximizing player will always win as +1", func() {
+               game := FakeGame{nextGames: []Game{
+                       FakeGame{isOver: true, winner: max}}}
+               Expect(Minimax(game, min)).To(Equal(1))
+       })
 })


### Step 8 -- Search

diff --git a/minimax/minimax.go b/minimax/minimax.go
index 34dadb6..84fd5fe 100644
--- a/minimax/minimax.go
+++ b/minimax/minimax.go
@@ -17,8 +17,16 @@ func Negamax(game Game, polarity int) int {
                return 0
        }

-       nextGame := game.NextGames()[0]
-       return -Negamax(nextGame, -1 * polarity)
+       bestScore := -100
+       for _, nextGame := range game.NextGames() {
+               scoreFromOpponentPerspective := Negamax(nextGame, -1 * polarity)
+               scoreFromMyPerspective := -1 * scoreFromOpponentPerspective
+               if scoreFromMyPerspective > bestScore {
+                       bestScore = scoreFromMyPerspective
+               }
+       }
+
+       return bestScore
 }
