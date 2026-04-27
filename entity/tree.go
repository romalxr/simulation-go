package entity

import (
	. "simulation/position"
)

type Tree struct {
	entityType EntityType
	position   Position
}

func NewTree(pos Position) *Tree {
	return &Tree{entityType: TypeTree, position: pos}
}

func (t *Tree) Type() EntityType {
	return t.entityType
}

func (t *Tree) Symbol() string {
	return "🌳"
}

func (t *Tree) Position() Position {
	return t.position
}
