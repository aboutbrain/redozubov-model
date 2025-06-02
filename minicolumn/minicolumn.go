package minicolumn

import (
	"github.com/aboutbrain/redozubov-model/config"
)

type Minicolumn struct {
	Neurons    []*Neuron
	Activated  bool
	Activation float64
}

func NewMinicolumn() *Minicolumn {
	mc := &Minicolumn{
		Neurons: make([]*Neuron, config.NumNeurons),
	}

	for i := range mc.Neurons {
		mc.Neurons[i] = NewNeuron()
	}

	return mc
}

func (mc *Minicolumn) ProcessPattern(input []complex128, context []float64) {
	maxActivation := 0.0
	mc.Activated = false

	for _, neuron := range mc.Neurons {
		activation := neuron.CalculateActivation(input, context)
		if activation > maxActivation {
			maxActivation = activation
		}
		if activation >= config.ActivationThreshold {
			mc.Activated = true
		}
	}

	mc.Activation = maxActivation
}
