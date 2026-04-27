package main

type Simulation struct {
	worldMap *WorldMap
	renderer *Renderer
}

func NewSimulation() *Simulation {
	wm := NewWorldMap(10, 10)
	wm.Generate(10)
	return &Simulation{
		worldMap: wm,
		renderer: NewRenderer(wm),
	}
}

func (s *Simulation) run() {
	s.renderer.Render()
}
