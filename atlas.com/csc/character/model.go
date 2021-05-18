package character

type Model struct {
	id    uint32
	hp    uint16
	maxHP uint16
}

func (a Model) Id() uint32 {
	return a.id
}

func (a Model) HP() uint16 {
	return a.hp
}

func (a Model) MaxHP() uint16 {
	return a.maxHP
}
