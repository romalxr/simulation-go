package entity

import (
	. "simulation/position"
)

type Rock struct {
	entityType EntityType
	position   Position
}

func NewRock(pos Position) *Rock {
	return &Rock{entityType: TypeRock, position: pos}
}

func (r *Rock) Type() EntityType {
	return r.entityType
}

func (r *Rock) Symbol() string {
	return "\U0001FAA8" // rock
}

func (r *Rock) Position() Position {
	return r.position
}

func (r *Rock) SetPosition(pos Position) {
	r.position = pos
}
