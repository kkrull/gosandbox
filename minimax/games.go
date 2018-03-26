package minimax

/* GameOver */

func GameOver() Game {
	return gameOver{}
}

type gameOver struct { }

func (gameOver) IsOver() bool {
	return true
}

/* QuickDraw */

func QuickDrawGame(winningSpace Space) Game {
	return quickDrawGame{}
}

type quickDrawGame struct {
}

func (quickDrawGame) IsOver() bool {
	return false
}

type Space interface {
}
