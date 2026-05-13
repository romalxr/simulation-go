package entity

import (
	. "simulation/position"
)

type Herbivore struct {
	entityType EntityType
	position   Position
	speed      int
	cooldown   int
	hitPoints  int
}

func NewHerbivore(pos Position, speed int, hp int) *Herbivore {
	return &Herbivore{entityType: TypeHerbivore, position: pos, speed: speed, hitPoints: hp}
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

func (h *Herbivore) Cooldown() int {
	return h.cooldown
}

func (h *Herbivore) SetCooldown(cd int) {
	h.cooldown = cd
}

func (h *Herbivore) Hp() int {
	return h.hitPoints
}

func (h *Herbivore) SetHp(hp int) {
	h.hitPoints = hp
}
