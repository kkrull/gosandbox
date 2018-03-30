package minimax

type CompletedGame struct {
	Winner string
}

func (game CompletedGame) FindWinner() string {
	return game.Winner
}

func (CompletedGame) IsOver() bool {
	return true
}

func (game CompletedGame) LegalMoves(player string) []Move {
	return []Move{}
}

type DuelGame struct {
	Space string
}

func (DuelGame) FindWinner() string {
	return ""
}

func (DuelGame) IsOver() bool {
	return false
}

func (game DuelGame) LegalMoves(player string) []Move {
	return []Move{
		Move{
			Game: CompletedGame{Winner: player},
			ClaimedSpace: game.Space,
		},
	}
}

