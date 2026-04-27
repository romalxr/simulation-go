package entity

import (
	. "simulation/position"
)

type Occupier interface {
	Type() EntityType
	Symbol() string
	Position() Position
}
