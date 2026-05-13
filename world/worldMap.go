package world

import (
	"errors"
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
	tiles  map[Position]Occupier
}

func NewWorldMap(width, height int) *WorldMap {
	return &WorldMap{
		tiles:  make(map[Position]Occupier),
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
	return wm.tiles[pos]
}

func (wm *WorldMap) SetTile(pos Position, occ Occupier) {
	wm.tiles[pos] = occ
}

func (wm *WorldMap) RemoveTile(pos Position) {
	delete(wm.tiles, pos)
}

func (wm *WorldMap) GetAll() []Occupier {
	res := make([]Occupier, 0, len(wm.tiles))

	for _, v := range wm.tiles {
		res = append(res, v)
	}

	return res
}

func (wm *WorldMap) IsEmpty(pos Position) bool {
	_, ok := wm.tiles[pos]
	return !ok
}

func (wm *WorldMap) IsValid(pos Position) bool {
	return pos.X >= 0 && pos.X < wm.width && pos.Y >= 0 && pos.Y < wm.height
}

func (wm *WorldMap) MoveTile(from, to Position) error {
	if wm.IsEmpty(from) {
		return errors.New("no tile at source position")
	}
	if !wm.IsEmpty(to) {
		return errors.New("target position is occupied")
	}
	occ, _ := wm.tiles[from]
	delete(wm.tiles, from)
	wm.tiles[to] = occ
	occ.SetPosition(to)
	return nil
}
