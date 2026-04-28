package main

import (
	"fmt"
	"simulation/position"
	"simulation/world"
)

type Renderer struct {
	worldMap world.WorldMapView
}

func NewRenderer(worldMap world.WorldMapView) *Renderer {
	return &Renderer{worldMap}
}

func (r *Renderer) Render(turn int) {
	// Clear screen
	fmt.Print("\033[2J\033[H")

	fmt.Printf("=== Simulation Turn %d ===\n", turn)

	r.printMap()
}

func (r *Renderer) printMap() {
	width := r.worldMap.GetWidth()
	height := r.worldMap.GetHeight()

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			occupier := r.worldMap.GetTile(position.Position{X: x, Y: y})
			icon := "|_"
			if occupier != nil {
				icon = occupier.Symbol()
			}
			fmt.Print(icon)
		}
		fmt.Println("")
	}
	fmt.Println("Enter 'p' to pause/resume or 'q' to exit")
}
