package minimax

/* ClosedSpace */

func ClosedSpace(id string) Space {
	return closedSpace{id: id}
}

type closedSpace struct {
	id string
}

func (closedSpace) Id() string {
	panic("implement me")
}

func (closedSpace) IsAvailable() bool {
	return false
}

/* OpenSpace */

func OpenSpace(id string) Space {
	return openSpace{id: id}
}

type openSpace struct {
	id string
}

func (space openSpace) Id() string {
	return space.id
}

func (openSpace) IsAvailable() bool {
	return true
}
