package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

var debugMode = false

func main() {
	if debugMode {
		f, _ := os.Create("simulation.log")
		defer f.Close()
		log.SetOutput(f)
	} else {
		log.SetOutput(io.Discard) // логи уходят в никуда
	}

	fmt.Println("===Simulation===")
	fmt.Println("")

	Simulation := NewSimulation()
	Simulation.Start()

	fmt.Println("")
	fmt.Println("===Bye, bye===")
	fmt.Println("")
}
