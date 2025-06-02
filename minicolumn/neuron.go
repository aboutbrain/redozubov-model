package minicolumn

import (
	"math"
	"math/rand"

	"github.com/aboutbrain/redozubov-model/config"
	"github.com/aboutbrain/redozubov-model/utils"
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
		n.Dendrites[i] = complex(rand.Float64()*0.5, math.Sin(phase))
	}

	for i := range n.Context {
		n.Context[i] = rand.NormFloat64() * 0.3
	}

	utils.ComplexNorm(n.Dendrites)

	return n
}

func (n *Neuron) CalculateActivation(input []complex128, context []float64) float64 {
	return ActivationFunction(input, context, n.Dendrites, n.Context)
}
