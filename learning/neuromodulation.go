package learning

import "github.com/aboutbrain/redozubov-model/minicolumn"

func ApplyDopamine(mc *minicolumn.Minicolumn, multiplier float64) {
	for _, neuron := range mc.Neurons {
		for i := range neuron.Dendrites {
			neuron.Dendrites[i] *= complex(multiplier, 0)
		}
	}
}
