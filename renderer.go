package main

import (
	"fmt"
	. "simulation/position"
)

type Renderer struct {
	worldMap WorldMapView
}

func NewRenderer(worldMap WorldMapView) *Renderer {
	return &Renderer{worldMap}
}

func (r *Renderer) Render() {
	// Clear screen
	fmt.Print("\033[2J\033[H")

	turn := 1
	fmt.Printf("=== Simulation Turn %d ===\n", turn)

	r.printMap()
}

func (r *Renderer) printMap() {
	width := r.worldMap.GetWidth()
	height := r.worldMap.GetHeight()

	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			occupier := r.worldMap.GetTile(Position{x, y})
			icon := "|_"
			if occupier != nil {
				icon = occupier.Symbol()
			}
			fmt.Print(icon)
		}
		fmt.Println("")
	}
}
