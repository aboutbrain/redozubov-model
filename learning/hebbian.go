package learning

import (
	"github.com/aboutbrain/redozubov-model/minicolumn"
)

// Убрана константа LearningRate, теперь передается как параметр
func ApplyHebbianLearning(neuron *minicolumn.Neuron, input []complex128, context []float64, learningRate float64) {
	for i := range neuron.Dendrites {
		for j, inVal := range input {
			if j < minicolumn.PatternSize && j < len(context) && j < len(neuron.Context) {
				contextFactor := 1 + context[j]*neuron.Context[j]
				modulated := inVal * complex(contextFactor, 0)

				neuron.Dendrites[i] += modulated * complex(learningRate, 0)
			}
		}
	}

	for j := range neuron.Context {
		if j < len(context) {
			neuron.Context[j] += learningRate * context[j]
		}
	}
}
