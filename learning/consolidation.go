package learning

import (
	"github.com/aboutbrain/redozubov-model/minicolumn"
	"math"
	"math/cmplx"
)

func Consolidate(neuron *minicolumn.Neuron) {
	// Более мягкая нормализация
	maxNorm := 0.0
	for i := range neuron.Dendrites {
		norm := cmplx.Abs(neuron.Dendrites[i])
		if norm > maxNorm {
			maxNorm = norm
		}
	}

	if maxNorm > 5.0 {
		scale := 5.0 / maxNorm
		for i := range neuron.Dendrites {
			neuron.Dendrites[i] *= complex(scale, 0)
		}
	}

	// Стабилизация контекстных весов
	for i := range neuron.Context {
		neuron.Context[i] = math.Tanh(neuron.Context[i])
	}
}
