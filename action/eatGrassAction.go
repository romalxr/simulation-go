package action

import (
	"math/rand"
	"simulation/entity"
	"simulation/position"
	"simulation/world"
)

type EatGrassAction struct {
}

func NewEatGrassAction() *EatGrassAction {
	return &EatGrassAction{}
}

func (a *EatGrassAction) Execute(wm *world.WorldMap) {
	directions := []position.Position{
		{X: 0, Y: -1}, // up
		{X: 0, Y: 1},  // down
		{X: -1, Y: 0}, // left
		{X: 1, Y: 0},  // right
	}

	for _, occ := range wm.GetAll() {
		herbivore, ok := occ.(*entity.Herbivore)
		if !ok {
			continue
		}

		if herbivore.Cooldown() > 0 {
			continue
		}

		pos := herbivore.Position()

		var available []position.Position
		for _, dir := range directions {
			newPos := position.Position{X: pos.X + dir.X, Y: pos.Y + dir.Y}
			if wm.IsValid(newPos) && !wm.IsEmpty(newPos) {
				neighbor := wm.GetTile(newPos)
				_, ok := neighbor.(*entity.Grass)
				if !ok {
					continue
				}
				available = append(available, newPos)
			}
		}

		if len(available) > 0 {
			randomIndex := rand.Intn(len(available))
			newPos := available[randomIndex]
			wm.RemoveTile(newPos)

			herbivore.SetCooldown(herbivore.Speed())
		}
	}
}
