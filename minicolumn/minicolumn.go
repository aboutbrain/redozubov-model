package minicolumn

import (
	"github.com/aboutbrain/redozubov-model/types"
	"math"
)

const (
	NumNeurons     = 12
	DendriteLength = 8
	PatternSize    = 5
)

type Minicolumn struct {
	Neurons     []*Neuron
	Activated   bool
	Activation  float64
	Astrocyte   types.AstrocyteInterface
	EnergyLevel float64
}

func NewMinicolumn(astro types.AstrocyteInterface) *Minicolumn {
	mc := &Minicolumn{
		Neurons:     make([]*Neuron, NumNeurons),
		Astrocyte:   astro,
		EnergyLevel: 1.0,
	}

	for i := range mc.Neurons {
		mc.Neurons[i] = NewNeuron()
	}

	return mc
}

func (mc *Minicolumn) ActivationThreshold() float64 {
	// Динамический порог в зависимости от энергии
	baseThreshold := 0.6
	energyFactor := 0.2 * (1.0 - mc.EnergyLevel)
	return math.Max(0.55, math.Min(0.85, baseThreshold+energyFactor))
}

func (mc *Minicolumn) ProcessPattern(input []complex128, context []float64) {
	// Получаем энергию от астроцита при критически низком уровне
	if mc.EnergyLevel < 0.2 && mc.Astrocyte != nil {
		mc.EnergyLevel += mc.Astrocyte.TransferEnergy()
	}

	// Кальциевая модуляция
	calciumMod := 1.0
	if mc.Astrocyte != nil {
		calciumMod = 1.0 + mc.Astrocyte.GetCalciumLevel()*0.6
	}

	maxActivation := 0.0
	mc.Activated = false
	threshold := mc.ActivationThreshold()

	for _, neuron := range mc.Neurons {
		activation := neuron.CalculateActivation(input, context) * calciumMod

		if activation > maxActivation {
			maxActivation = activation
		}
		if activation >= threshold {
			mc.Activated = true
		}
	}

	mc.Activation = maxActivation

	// Расходуем энергию только при активации
	if mc.Activated {
		mc.EnergyLevel -= 0.04 // Уменьшили расход
	} else {
		mc.EnergyLevel -= 0.005 // Минимальный расход в покое
	}

	if mc.EnergyLevel < 0 {
		mc.EnergyLevel = 0
	}
}

func (mc *Minicolumn) Rest() {
	// Усиленное восстановление с учетом астроцитарной поддержки
	recovery := 0.5
	if mc.Astrocyte != nil {
		recovery += 0.3 * mc.Astrocyte.GetCalciumLevel()
	}
	mc.EnergyLevel = math.Min(1.0, mc.EnergyLevel+recovery)
}
