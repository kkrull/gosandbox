package minimax

/* GameOver */

func GameOver() Game {
	return gameOver{}
}

type gameOver struct { }

func (gameOver) OpenSpaces() []Space {
	return nil
}

func (gameOver) IsOver() bool {
	return true
}

/* QuickDraw */

func QuickDrawGame(winningSpace Space) Game {
	return quickDrawGame{winningSpace: winningSpace}
}

type quickDrawGame struct {
	winningSpace Space
}

func (game quickDrawGame) OpenSpaces() []Space {
	spaces := make([]Space, 0)
	spaces = append(spaces, game.winningSpace)
	return spaces
}

func (quickDrawGame) IsOver() bool {
	return false
}

/* CaptureTheFlag */


func CaptureTheFlagGame() Game {
	return captureTheFlagGame{
		maxBase: "MaxBase",
		minBase: "MinBase",
	}
}

type captureTheFlagGame struct {
	maxBase Space
	minBase Space
}

func (game captureTheFlagGame) OpenSpaces() []Space {
	return []Space {
		game.maxBase,
		game.minBase,
	}
}

func (game captureTheFlagGame) IsOver() bool {
	return len(game.OpenSpaces()) == 0
}
