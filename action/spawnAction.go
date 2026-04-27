package action

import (
	"math/rand"
	"simulation/entity"
	"simulation/position"
	"simulation/world"
)

type SpawnAction struct {
	factory func(position.Position) entity.Occupier
	count   int
}

func NewSpawnAction(factory func(position.Position) entity.Occupier, count int) *SpawnAction {
	return &SpawnAction{factory: factory, count: count}
}

func (a *SpawnAction) Execute(wm *world.WorldMap) {
	width := wm.GetWidth()
	height := wm.GetHeight()
	totalCells := width * height
	if a.count > totalCells {
		a.count = totalCells
	}

	placed := 0
	for placed < a.count {
		x := rand.Intn(width)
		y := rand.Intn(height)
		pos := position.Position{X: x, Y: y}
		if wm.GetTile(pos) == nil {
			wm.SetTile(pos, a.factory(pos))
			placed++
		}
	}
}
