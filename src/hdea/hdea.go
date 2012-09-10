package hdea

import (
	"optprobs"
)

type HdEA struct {
	// The problem to optimize
	problem optprobs.FunctionEvaluator
	// Number of dimensions
	dimension int32
	// Size of maintained population
	populationSize int32

}

// New returns a new instance
func NewHdEA(problem optprobs.FunctionEvaluator, dimension, populationSize int32) (*HdEA) {
	hdea := new(HdEA)
	hdea.problem = problem
	hdea.dimension = dimension
	hdea.populationSize = populationSize
	return hdea
}

func (*HdEA) Name() string {
	return "HdEA"
}

func (*HdEA) Optimize() (*Individuum) {
	return nil
}
