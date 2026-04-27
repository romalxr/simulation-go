package main

import (
	"simulation/action"
	"simulation/entity"
	"simulation/position"
	"simulation/world"
	"time"
)

type Simulation struct {
	worldMap    *world.WorldMap
	renderer    *Renderer
	initActions []action.Action
}

func NewSimulation() *Simulation {

	wm := world.NewWorldMap(10, 10)

	initActions := []action.Action{
		action.NewSpawnAction(func(pos position.Position) entity.Occupier {
			return entity.NewGrass(pos)
		}, 10),
		action.NewSpawnAction(func(pos position.Position) entity.Occupier {
			return entity.NewTree(pos)
		}, 3),
		action.NewSpawnAction(func(pos position.Position) entity.Occupier {
			return entity.NewRock(pos)
		}, 2),
	}

	return &Simulation{
		worldMap:    wm,
		renderer:    NewRenderer(wm),
		initActions: initActions,
	}
}

func (s *Simulation) Start() {

	for _, action := range s.initActions {
		action.Execute(s.worldMap)
	}

	for i := 0; i < 10; i++ {
		s.renderer.Render()
		time.Sleep(200 * time.Millisecond)
	}
}
