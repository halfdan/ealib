package hdea

import (
	"optprobs"
)

type HdEA struct {
	problem optprobs.FunctionEvaluator
}

// New returns a new instance
func New(problem optprobs.FunctionEvaluator) (*HdEA) {
	hdea := new(HdEA)
	hdea.problem = problem
	return hdea
}

func (*HdEA) Name() string {
	return "HdEA"
}
