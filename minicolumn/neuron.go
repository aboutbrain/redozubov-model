package minicolumn

import (
	"math"
	"math/cmplx"
	"math/rand"
)

type Neuron struct {
	Dendrites []complex128
	Context   []float64
}

func NewNeuron() *Neuron {
	n := &Neuron{
		Dendrites: make([]complex128, DendriteLength),
		Context:   make([]float64, PatternSize),
	}

	for i := range n.Dendrites {
		phase := rand.Float64() * 2 * math.Pi
		n.Dendrites[i] = complex(rand.Float64()*0.5, math.Sin(phase))
	}

	for i := range n.Context {
		n.Context[i] = rand.NormFloat64() * 0.3
	}

	return n
}

func (n *Neuron) CalculateActivation(input []complex128, context []float64) float64 {
	if len(input) == 0 {
		return 0.0
	}

	total := 0.0 + 0i

	// Обрабатываем каждый элемент входного паттерна
	for i := 0; i < len(input) && i < PatternSize; i++ {
		// Применяем контекстную модуляцию
		contextFactor := 1.0
		if i < len(context) && i < len(n.Context) {
			contextFactor = 1 + context[i]*n.Context[i]
		}

		modulated := input[i] * complex(contextFactor, 0)

		// Обрабатываем каждую дендритную ветвь
		for j := 0; j < DendriteLength && j < len(n.Dendrites); j++ {
			total += modulated * cmplx.Conj(n.Dendrites[j])
		}
	}

	// Нормализуем и вычисляем активацию
	magnitude := cmplx.Abs(total) / float64(len(input)*DendriteLength)
	return math.Pow(magnitude, 2)
}
