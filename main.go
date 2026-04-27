package main

import "fmt"

func main() {
	fmt.Println("===Simulation===")
	fmt.Println("")

	Simulation := NewSimulation()
	Simulation.run()
}
