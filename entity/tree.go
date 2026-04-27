package entity

type Tree struct {
	entityType EntityType
}

func NewTree() *Tree {
	return &Tree{entityType: TypeTree}
}

func (g *Tree) Type() EntityType {
	return g.entityType
}

func (g *Tree) Symbol() string {
	return "🌳"
}
