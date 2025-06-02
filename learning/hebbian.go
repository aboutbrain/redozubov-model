package learning

import (
	"github.com/aboutbrain/redozubov-model/minicolumn"
)

const LearningRate = 0.1

func ApplyHebbianLearning(neuron *minicolumn.Neuron, input []complex128, context []float64) {
	for i := range neuron.Dendrites {
		for j, inVal := range input {
			// Проверяем границы массивов
			if j < len(context) && j < len(neuron.Context) {
				modulated := inVal * complex(1+context[j]*neuron.Context[j], 0)
				neuron.Dendrites[i] += modulated * complex(LearningRate, 0)
			}
		}
	}

	for j := range neuron.Context {
		if j < len(context) {
			neuron.Context[j] += LearningRate * context[j]
		}
	}
}
