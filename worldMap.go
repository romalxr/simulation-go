package main

import (
	"math/rand"
	. "simulation/entity"
	. "simulation/position"
)

type WorldMapView interface {
	GetWidth() int
	GetHeight() int
	GetTile(pos Position) Occupier
}

type WorldMap struct {
	width  int
	height int
	data   map[Position]Occupier
}

func (wm *WorldMap) GetWidth() int  { return wm.width }
func (wm *WorldMap) GetHeight() int { return wm.height }

func NewWorldMap(width, height int) *WorldMap {
	return &WorldMap{
		data:   make(map[Position]Occupier),
		width:  width,
		height: height,
	}
}

func (wm *WorldMap) Generate(tile Occupier, numGrass int) {
	totalCells := wm.width * wm.height
	if numGrass > totalCells {
		numGrass = totalCells
	}

	placed := 0
	for placed < numGrass {
		x := rand.Intn(wm.width)
		y := rand.Intn(wm.height)
		pos := Position{X: x, Y: y}
		if _, exists := wm.data[pos]; !exists {
			wm.data[pos] = tile.MoveTo(pos)
			placed++
		}
	}
}

func (wm *WorldMap) GetTile(pos Position) Occupier {
	return wm.data[pos]
}
