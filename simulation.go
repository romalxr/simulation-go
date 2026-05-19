package main

import (
	"fmt"
	"log"
	"math/rand"
	"simulation/action"
	"simulation/entity"
	"simulation/position"
	"simulation/world"
	"sync/atomic"
	"time"
)

type Simulation struct {
	worldMap    *world.WorldMap
	renderer    *Renderer
	initActions []action.Action
	turnActions []action.Action
	isPaused    atomic.Bool
	turn        int
	isStopped   atomic.Bool
}

func NewSimulation() *Simulation {

	wm := world.NewWorldMap(10, 20)

	initActions := []action.Action{
		action.NewSpawnAction(func(pos position.Position) entity.Occupier {
			return entity.NewGrass(pos)
		}, 20),
		action.NewSpawnAction(func(pos position.Position) entity.Occupier {
			return entity.NewTree(pos)
		}, 3),
		action.NewSpawnAction(func(pos position.Position) entity.Occupier {
			return entity.NewRock(pos)
		}, 2),
		action.NewSpawnAction(func(pos position.Position) entity.Occupier {
			speed := 2 + rand.Intn(4)
			return entity.NewHerbivore(pos, speed, 50)
		}, 5),
		action.NewSpawnAction(func(pos position.Position) entity.Occupier {
			speed := 1 + rand.Intn(4)
			return entity.NewPredator(pos, speed, 70, 20)
		}, 3),
	}

	turnActions := []action.Action{
		action.NewUpdateCooldownAction(),
		action.NewEatGrassAction(),
		action.NewAttackAction(),
		action.NewSeekMoveAction(),
		action.NewRandomMoveAction(),
		action.NewRespawnAction(10, 5),
	}

	return &Simulation{
		worldMap:    wm,
		renderer:    NewRenderer(wm),
		initActions: initActions,
		turnActions: turnActions,
	}
}

func (s *Simulation) Start() {

	for _, action := range s.initActions {
		action.Execute(s.worldMap)
	}

	go s.handleCommands()

	s.turn = 0
	s.isPaused.Store(false)

	for {
		if s.isStopped.Load() {
			break
		}
		if !s.isPaused.Load() {
			s.NextTurn()
		}
		time.Sleep(500 * time.Millisecond)
	}
}

func (s *Simulation) NextTurn() {

	log.Printf("Simulation: Выполняется ход %d", s.turn)
	for _, act := range s.turnActions {
		act.Execute(s.worldMap)
	}
	s.turn++
	s.renderer.Render(s.turn)
}

func (s *Simulation) Pause() {
	s.isPaused.Store(!s.isPaused.Load())
}

func (s *Simulation) Stop() {
	s.isStopped.Store(true)
}

func (s *Simulation) handleCommands() {
	var cmd string

	for {
		fmt.Scanln(&cmd)

		switch cmd {
		case "p":
			s.Pause()
		case "q":
			s.Stop()
		}
	}
}
