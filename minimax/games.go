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
