package learning

import (
	"github.com/aboutbrain/redozubov-model/minicolumn"
	"github.com/aboutbrain/redozubov-model/utils"
)

const LearningRate = 0.1

func ApplyHebbianLearning(neuron *minicolumn.Neuron, input []complex128, context []float64) {
	for i := range neuron.Dendrites {
		for j, inVal := range input {
			if j >= len(neuron.Context) || j >= len(context) {
				continue
			}
			modulation := 1.0 + context[j]*neuron.Context[j]
			modulated := inVal * complex(modulation, 0)
			neuron.Dendrites[i] += modulated * complex(LearningRate, 0)
		}
	}

	for j := range neuron.Context {
		if j < len(context) {
			neuron.Context[j] += LearningRate * context[j]
		}
	}

	utils.ComplexNorm(neuron.Dendrites)
}
