package main

import "fmt"

func main() {
	fmt.Println("===Simulation===")
	fmt.Println("")

	Simulation := NewSimulation()
	Simulation.Start()

	fmt.Println("")
	fmt.Println("===Bye, bye===")
	fmt.Println("")
}
