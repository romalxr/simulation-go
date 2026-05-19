package action

import (
	"log"
	"simulation/entity"
	pathfinding "simulation/pathFinding"
	"simulation/world"
)

type SeekMoveAction struct {
}

func NewSeekMoveAction() *SeekMoveAction {
	return &SeekMoveAction{}
}

func (a *SeekMoveAction) Execute(wm *world.WorldMap) {

	for _, occ := range wm.GetAll() {
		creature, ok := occ.(entity.Creature)
		if !ok {
			continue
		}

		if creature.Cooldown() > 0 {
			log.Printf("SeekMoveAction: Существо %s %p на кулдауне %d", creature.Type(), creature, creature.Cooldown())
			continue
		}

		var targetType entity.EntityType
		if creature.Type() == entity.TypePredator {
			targetType = entity.TypeHerbivore
		} else if creature.Type() == entity.TypeHerbivore {
			targetType = entity.TypeGrass
		} else {
			continue
		}

		oldPos := creature.Position()
		nextStep := pathfinding.NextStepToTarget(wm, oldPos, targetType, 5)
		if nextStep != nil {
			log.Printf("SeekMoveAction: Существо %s %p нашли след шаг", creature.Type(), creature)
			wm.MoveTile(oldPos, *nextStep)
			creature.SetCooldown(creature.Speed())
		} else {
			log.Printf("SeekMoveAction: Существо %s %p не нашли след шаг", creature.Type(), creature)
		}

	}
}
