package optprobs

import "math"

type AckleyFunction struct{}

func (a AckleyFunction) Evaluate(solution []float64) float64 {
	sum1 := 0.0
	sum2 := 0.0
	dim := len(solution)

	for _, v := range solution {
		sum1 += v * v
		sum2 += math.Cos(2 * math.Pi * v)
	}

	return -20.0*math.Exp(-0.2*math.Sqrt(sum1/float64(dim))) - math.Exp(sum2/float64(dim)) + 20.0 + math.E
}

func (a AckleyFunction) Name() string {
	return "Ackley Function"
}
