package action

import (
	"math/rand"
	"simulation/entity"
	"simulation/position"
	"simulation/world"
)

type AttackAction struct {
}

func NewAttackAction() *AttackAction {
	return &AttackAction{}
}

func (a *AttackAction) Execute(wm *world.WorldMap) {
	directions := []position.Position{
		{X: 0, Y: -1}, // up
		{X: 0, Y: 1},  // down
		{X: -1, Y: 0}, // left
		{X: 1, Y: 0},  // right
	}

	for _, occ := range wm.GetAll() {
		predator, ok := occ.(*entity.Predator)
		if !ok {
			continue
		}

		if predator.Cooldown() > 0 {
			continue
		}

		pos := predator.Position()

		var available []position.Position
		for _, dir := range directions {
			newPos := position.Position{X: pos.X + dir.X, Y: pos.Y + dir.Y}
			if wm.IsValid(newPos) && !wm.IsEmpty(newPos) {
				neighbor := wm.GetTile(newPos)
				_, ok := neighbor.(*entity.Herbivore)
				if !ok {
					continue
				}
				available = append(available, newPos)
			}
		}

		if len(available) > 0 {
			randomIndex := rand.Intn(len(available))
			newPos := available[randomIndex]

			neighbor := wm.GetTile(newPos)
			herbivore, ok := neighbor.(*entity.Herbivore)
			if !ok {
				continue
			}

			herbivore.SetHp(herbivore.Hp() - predator.Ap())
			if herbivore.Hp() <= 0 {
				wm.RemoveTile(newPos)
			}

			predator.SetCooldown(herbivore.Speed())
		}
	}
}
