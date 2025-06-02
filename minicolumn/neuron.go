package minicolumn

import (
	"math"
	"math/rand"

	"github.com/aboutbrain/redozubov-model/config"
)

type Neuron struct {
	Dendrites []complex128
	Context   []float64
}

func NewNeuron() *Neuron {
	n := &Neuron{
		Dendrites: make([]complex128, config.DendriteLength),
		Context:   make([]float64, config.PatternSize),
	}

	for i := range n.Dendrites {
		phase := rand.Float64() * 2 * math.Pi
		amplitude := 0.8 + rand.Float64()*0.4
		n.Dendrites[i] = complex(
			amplitude*math.Cos(phase),
			amplitude*math.Sin(phase),
		)
	}

	for i := range n.Context {
		n.Context[i] = 0.5*rand.NormFloat64() + 0.1
	}

	return n
}

// Добавляем отсутствующий метод CalculateActivation
func (n *Neuron) CalculateActivation(input []complex128, context []float64) float64 {
	return ActivationFunction(input, context, n.Dendrites, n.Context)
}
