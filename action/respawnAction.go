package action

import (
	"math/rand"
	"simulation/entity"
	"simulation/world"
)

type RespawnAction struct {
	herbivoreСhance int
	grassСhance     int
}

func NewRespawnAction(grassСhance int, herbivoreСhance int) *RespawnAction {
	return &RespawnAction{grassСhance: grassСhance, herbivoreСhance: herbivoreСhance}
}

func (a *RespawnAction) Execute(wm *world.WorldMap) {
	if rand.Intn(100) < a.grassСhance {

		if pos := wm.GetRandomEmptyPos(); pos != nil {
			wm.SetTile(*pos, entity.NewGrass(*pos))
		}
	}

	if rand.Intn(100) < a.herbivoreСhance {

		if pos := wm.GetRandomEmptyPos(); pos != nil {
			speed := 2 + rand.Intn(4)
			wm.SetTile(*pos, entity.NewHerbivore(*pos, speed, 50))
		}
	}
}
