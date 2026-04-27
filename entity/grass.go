package entity

import (
	. "simulation/position"
)

type Grass struct {
	entityType EntityType
	position   Position
}

func NewGrass(pos Position) *Grass {
	return &Grass{entityType: TypeGrass, position: pos}
}

func (g *Grass) Type() EntityType {
	return g.entityType
}

func (g *Grass) Symbol() string {
	return "🌿"
}

func (g *Grass) Position() Position {
	return g.position
}
