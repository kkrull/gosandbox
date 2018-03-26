package minimax


func GameOver() Game {
	return gameOver{}
}

type gameOver struct { }

func (gameOver) IsOver() bool {
	return true
}
