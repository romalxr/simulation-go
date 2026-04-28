package entity

type Creature interface {
	Occupier
	Speed() int
}
