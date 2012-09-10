package main

import (
	"fmt"
	"optprobs"
	"hdea"
)

func main() {
	fmt.Printf("ealib v0.1\n")
	ackley, _ := optprobs.GetEvaluator(0)
	fmt.Printf("Problem: %s\n", ackley.Name())

	// New HdEA with dimension=30, populationSize=100
	ea := hdea.NewHdEA(ackley, 30, 100)
	fmt.Printf("%s", ea.Name())

}
