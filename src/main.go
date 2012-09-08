package main

import (
	"fmt"
	"hdea"
)

func main() {
	fmt.Printf("ealib v0.1")
	ea := hdea.New()

	fmt.Printf(ea.Name())
}
