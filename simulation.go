package main

import (
	. "simulation/entity"
	"time"
)

type Simulation struct {
	worldMap *WorldMap
	renderer *Renderer
}

func NewSimulation() *Simulation {
	wm := NewWorldMap(10, 10)
	wm.Generate(NewGrass(), 10)
	wm.Generate(NewTree(), 3)
	wm.Generate(NewRock(), 2)
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
