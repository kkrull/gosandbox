package minimax

/* ClosedSpace */

type ClosedSpace struct {
	IdValue string
}

func (ClosedSpace) Id() string {
	panic("implement me")
}

func (ClosedSpace) IsAvailable() bool {
	return false
}

/* OpenSpace */

type OpenSpace struct {
	IdValue string
}

func (space OpenSpace) Id() string {
	return space.IdValue
}

func (OpenSpace) IsAvailable() bool {
	return true
}
