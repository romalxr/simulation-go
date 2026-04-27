package entity

type Rock struct {
	entityType EntityType
}

func NewRock() *Rock {
	return &Rock{entityType: TypeRock}
}

func (g *Rock) Type() EntityType {
	return g.entityType
}

func (g *Rock) Symbol() string {
	return "🪨"
}
