package entity

import (
	. "simulation/position"
)

type Herbivore struct {
	entityType EntityType
	position   Position
	speed      int
}

func NewHerbivore(pos Position, speed int) *Herbivore {
	return &Herbivore{entityType: TypeHerbivore, position: pos, speed: speed}
}

func (h *Herbivore) Type() EntityType {
	return h.entityType
}

func (h *Herbivore) Symbol() string {
	return "\U0001F40F" // sheep
}

func (h *Herbivore) Position() Position {
	return h.position
}

func (h *Herbivore) SetPosition(pos Position) {
	h.position = pos
}

func (h *Herbivore) Speed() int {
	return h.speed
}
