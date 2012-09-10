package hdea

type Individuum struct {
	Data []float64
	Fitness float64
}

func NewIndividuum(data []float64) (*Individuum) {
	indi := new(Individuum)
	indi.Data = data
	return indi
}

func (indi *Individuum) SetFitness(fitness float64) {
	indi.Fitness = fitness
}

