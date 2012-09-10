package optprobs

type RosenbrockFunction struct{}

func (a RosenbrockFunction) Evaluate(solution []float64) float64 {
	sum := 0.0
	d := make([]float64, len(solution))

	for i, v := range solution {
		d[i] = v + float64(1)
	}

	for i := 0; i < len(solution)-1; i++ {
		temp1 := (d[i] * d[i]) - d[i+1]
		temp2 := d[i] - 1.0
		sum += (100.0 * temp1 * temp1) + (temp2 * temp2)
	}

	return sum
}

func (a RosenbrockFunction) Name() string {
	return "Rosenbrock Function"
}
