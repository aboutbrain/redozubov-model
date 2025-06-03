package learning

import (
	"github.com/aboutbrain/redozubov-model/minicolumn"
	"math/cmplx"
)

func Consolidate(neuron *minicolumn.Neuron) {
	// Смягчаем нормализацию
	totalNorm := 0.0
	for i := range neuron.Dendrites {
		norm := cmplx.Abs(neuron.Dendrites[i])
		totalNorm += norm
	}

	if totalNorm > 10.0 { // Более мягкое ограничение
		scale := 10.0 / totalNorm
		for i := range neuron.Dendrites {
			neuron.Dendrites[i] *= complex(scale, 0)
		}
	}

	// Нормализация контекстных весов
	for i := range neuron.Context {
		if neuron.Context[i] > 1.0 {
			neuron.Context[i] = 1.0
		} else if neuron.Context[i] < -1.0 {
			neuron.Context[i] = -1.0
		}
	}
}
