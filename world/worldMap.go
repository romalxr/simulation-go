package world

import (
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

func NewWorldMap(width, height int) *WorldMap {
	return &WorldMap{
		data:   make(map[Position]Occupier),
		width:  width,
		height: height,
	}
}

func (wm *WorldMap) GetWidth() int {
	return wm.width
}

func (wm *WorldMap) GetHeight() int {
	return wm.height
}

func (wm *WorldMap) GetTile(pos Position) Occupier {
	return wm.data[pos]
}

func (wm *WorldMap) SetTile(pos Position, occ Occupier) {
	wm.data[pos] = occ
}
