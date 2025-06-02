package learning

import (
	"github.com/aboutbrain/redozubov-model/config"
	"github.com/aboutbrain/redozubov-model/minicolumn"
)

func ApplyHebbianLearning(neuron *minicolumn.Neuron, input []complex128, context []float64) {
	for i := range neuron.Dendrites {
		// Упрощенное правило Хебба
		neuron.Dendrites[i] += input[i] * complex(config.LearningRate, 0)
	}

	for j := range neuron.Context {
		// Сильнее реагируем на контекст
		neuron.Context[j] += 2.0 * config.LearningRate * context[j]
	}
}
