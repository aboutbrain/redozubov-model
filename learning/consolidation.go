package learning

import (
	"github.com/aboutbrain/redozubov-model/minicolumn"
	"math/cmplx"
)

func Consolidate(neuron *minicolumn.Neuron) {
	for i := range neuron.Dendrites {
		norm := cmplx.Abs(neuron.Dendrites[i])
		if norm > 1.0 {
			neuron.Dendrites[i] /= complex(norm, 0)
		}
	}

	for i := range neuron.Context {
		if neuron.Context[i] > 1.0 {
			neuron.Context[i] = 1.0
		} else if neuron.Context[i] < -1.0 {
			neuron.Context[i] = -1.0
		}
	}
}
