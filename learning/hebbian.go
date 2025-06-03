package learning

import (
	"github.com/aboutbrain/redozubov-model/minicolumn"
)

const LearningRate = 0.5 // Увеличили скорость обучения

func ApplyHebbianLearning(neuron *minicolumn.Neuron, input []complex128, context []float64) {
	for i := range neuron.Dendrites {
		for j, inVal := range input {
			if j < minicolumn.PatternSize && j < len(context) && j < len(neuron.Context) {
				contextFactor := 1 + context[j]*neuron.Context[j]
				modulated := inVal * complex(contextFactor, 0)

				// Уменьшаем влияние обучения
				neuron.Dendrites[i] += modulated * complex(LearningRate*0.5, 0)
			}
		}
	}

	for j := range neuron.Context {
		if j < len(context) {
			// Более агрессивное обновление контекста
			neuron.Context[j] += 2 * LearningRate * context[j]
		}
	}
}
