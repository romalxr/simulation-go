package pathfinding

import (
	"simulation/entity"
	"simulation/position"
	"simulation/world"
)

func NextStepToTarget(wm *world.WorldMap, start position.Position, targetType entity.EntityType, maxDepth int) *position.Position {
	directions := []position.Position{
		{X: 0, Y: -1}, {X: 0, Y: 1}, {X: -1, Y: 0}, {X: 1, Y: 0},
	}

	currentLevel := []position.Position{start}
	visited := make(map[position.Position]bool)
	firstStep := make(map[position.Position]position.Position)
	visited[start] = true

	for depth := 0; depth < maxDepth && len(currentLevel) > 0; depth++ {
		nextLevel := []position.Position{}

		for _, current := range currentLevel {
			if !wm.IsEmpty(current) && wm.GetTile(current).Type() == targetType {
				if step, ok := firstStep[current]; ok {
					return &step
				}
				return nil
			}

			for _, dir := range directions {
				next := position.Position{X: current.X + dir.X, Y: current.Y + dir.Y}

				if !wm.IsValid(next) {
					continue
				}
				if visited[next] {
					continue
				}

				if !wm.IsEmpty(next) && wm.GetTile(next).Type() != targetType {
					continue
				}

				visited[next] = true
				if current == start {
					firstStep[next] = next
				} else {
					firstStep[next] = firstStep[current]
				}
				nextLevel = append(nextLevel, next)
			}
		}
		currentLevel = nextLevel
	}

	return nil
}
