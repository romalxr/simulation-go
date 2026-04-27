package entity

type Grass struct {
	entityType EntityType
}

func NewGrass() *Grass {
	return &Grass{entityType: TypeGrass}
}

func (g *Grass) Type() EntityType {
	return g.entityType
}

func (g *Grass) Symbol() string {
	return "🌿"
}
