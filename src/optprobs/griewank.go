package optprobs

import "math"

type GriewankFunction struct{}

func (a GriewankFunction) Evaluate(solution []float64) float64 {
	sum := 0.0
	product := 1.0

	for i, v := range solution {
		sum += ((v * v) / 4000.0)
		product *= math.Cos(v / math.Sqrt(float64(i+1)))
	}
	return (sum - product + 1.0)
}

func (a GriewankFunction) Name() string {
	return "Griewank Function"
}
