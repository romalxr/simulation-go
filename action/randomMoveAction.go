package action

import (
	"math/rand"
	"simulation/entity"
	"simulation/position"
	"simulation/world"
)

type RandomMoveAction struct {
}

func NewRandomMoveAction() *RandomMoveAction {
	return &RandomMoveAction{}
}

func (a *RandomMoveAction) Execute(wm *world.WorldMap) {
	directions := []position.Position{
		{X: 0, Y: -1}, // up
		{X: 0, Y: 1},  // down
		{X: -1, Y: 0}, // left
		{X: 1, Y: 0},  // right
	}

	for _, occ := range wm.GetAll() {
		creature, ok := occ.(entity.Creature)
		if !ok {
			continue
		}

		oldPos := creature.Position()

		var available []position.Position
		for _, dir := range directions {
			newPos := position.Position{X: oldPos.X + dir.X, Y: oldPos.Y + dir.Y}
			if wm.IsValid(newPos) && wm.IsEmpty(newPos) {
				available = append(available, newPos)
			}
		}

		if len(available) > 0 {
			randomIndex := rand.Intn(len(available))
			newPos := available[randomIndex]
			wm.MoveTile(oldPos, newPos)
		}
	}
}
