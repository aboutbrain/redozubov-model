package minicolumn

import "math/cmplx"

func ActivationFunction(
	input []complex128,
	context []float64,
	dendrites []complex128,
	neuronContext []float64,
) float64 {
	total := 0.0 + 0i

	// Упрощаем вычисления - используем только первые dendriteLength входов
	for i := 0; i < len(dendrites) && i < len(input); i++ {
		modulation := 3.0 + 5.0*context[i]*neuronContext[i]
		modulated := input[i] * complex(modulation, 0)
		total += modulated * cmplx.Conj(dendrites[i])
	}

	return cmplx.Abs(total) / float64(len(dendrites))
}
