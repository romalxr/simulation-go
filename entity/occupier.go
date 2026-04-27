package entity

type Occupier interface {
	Type() EntityType
	Symbol() string
}
