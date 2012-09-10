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

	ea := hdea.New(ackley)

}
