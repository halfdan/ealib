package optprobs

type Evaluator interface {
	Evaluate([]float64) float64
}
