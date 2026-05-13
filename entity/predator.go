package entity

import (
	. "simulation/position"
)

type Predator struct {
	entityType  EntityType
	position    Position
	speed       int
	cooldown    int
	hitPoints   int
	attackPower int
}

func NewPredator(pos Position, speed int, hp int, ap int) *Predator {
	return &Predator{
		entityType:  TypePredator,
		position:    pos,
		speed:       speed,
		hitPoints:   hp,
		attackPower: ap,
	}
}

func (p *Predator) Type() EntityType {
	return p.entityType
}

func (p *Predator) Symbol() string {
	return "\U0001F43A" // wolf
}

func (p *Predator) Position() Position {
	return p.position
}

func (p *Predator) SetPosition(pos Position) {
	p.position = pos
}

func (p *Predator) Speed() int {
	return p.speed
}

func (p *Predator) Cooldown() int {
	return p.cooldown
}

func (p *Predator) SetCooldown(cd int) {
	p.cooldown = cd
}

func (p *Predator) Hp() int {
	return p.hitPoints
}

func (p *Predator) SetHp(hp int) {
	p.hitPoints = hp
}

func (p *Predator) Ap() int {
	return p.attackPower
}
