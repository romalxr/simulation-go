package main

import (
	. "simulation/entity"
	. "simulation/position"
	"time"
)

type Simulation struct {
	worldMap *WorldMap
	renderer *Renderer
}

func NewSimulation() *Simulation {

	wm := NewWorldMap(10, 10)
	wm.Generate(NewGrass(Position{}), 10)
	wm.Generate(NewTree(Position{}), 3)
	wm.Generate(NewRock(Position{}), 2)
	return &Simulation{
		worldMap: wm,
		renderer: NewRenderer(wm),
	}
}

func (s *Simulation) run() {
	for i := 0; i < 10; i++ {
		s.renderer.Render()
		time.Sleep(200 * time.Millisecond)
	}
}
