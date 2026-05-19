package action

import (
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
		if pos := wm.GetRandomEmptyPos(); pos != nil {
			wm.SetTile(*pos, a.factory(*pos))
			placed++
		}
	}
}
