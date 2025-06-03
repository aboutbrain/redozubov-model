package minicolumn

import (
	"github.com/aboutbrain/redozubov-model/types"
)

const (
	NumNeurons          = 100  // Увеличили с 12 до 100
	ActivationThreshold = 0.85 // Возвращаем оригинальный порог
	DendriteLength      = 50   // Увеличили с 8 до 50
	PatternSize         = 5
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

func (mc *Minicolumn) ProcessPattern(input []complex128, context []float64) {
	// Проверяем наличие астроцита и получаем энергию
	if mc.Astrocyte != nil {
		if mc.EnergyLevel < 0.3 {
			mc.EnergyLevel += mc.Astrocyte.TransferEnergy()
		}
	}

	// Вычисляем кальциевую модуляцию
	calciumMod := 1.0
	if mc.Astrocyte != nil {
		calciumMod = 1.0 + mc.Astrocyte.GetCalciumLevel()*0.5
	}

	maxActivation := 0.0
	mc.Activated = false

	// Обрабатываем каждый нейрон в миниколонке
	for _, neuron := range mc.Neurons {
		activation := neuron.CalculateActivation(input, context) * calciumMod

		if activation > maxActivation {
			maxActivation = activation
		}
		if activation >= ActivationThreshold {
			mc.Activated = true
		}
	}

	mc.Activation = maxActivation

	// Расходуем энергию на обработку
	if mc.Astrocyte != nil {
		mc.EnergyLevel -= 0.05
		if mc.EnergyLevel < 0 {
			mc.EnergyLevel = 0
		}
	}

	// Увеличиваем расход энергии при активации
	if mc.Activated {
		mc.EnergyLevel -= 0.15
	} else {
		mc.EnergyLevel -= 0.02
	}

	if mc.EnergyLevel < 0 {
		mc.EnergyLevel = 0
	}
}
