package action

import (
	"simulation/entity"
	"simulation/world"
)

type UpdateCooldownAction struct {
}

func NewUpdateCooldownAction() *UpdateCooldownAction {
	return &UpdateCooldownAction{}
}

func (a *UpdateCooldownAction) Execute(wm *world.WorldMap) {
	for _, occ := range wm.GetAll() {
		creature, ok := occ.(entity.Creature)
		if !ok {
			continue
		}

		creature.SetCooldown(max(creature.Cooldown()-1, 0))
	}
}
