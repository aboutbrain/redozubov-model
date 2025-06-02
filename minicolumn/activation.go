package minicolumn

import "math/cmplx"

// ActivationFunction вычисляет уровень активации на основе волновой интерференции
func ActivationFunction(
	input []complex128,
	context []float64,
	dendrites []complex128,
	neuronContext []float64,
) float64 {
	total := 0.0 + 0i

	for i := range input {
		// Модуляция входного сигнала контекстом
		modulated := input[i] * complex(1+context[i]*neuronContext[i], 0)

		// Интерференция с дендритными весами
		for j := 0; j < len(dendrites); j++ {
			total += modulated * cmplx.Conj(dendrites[j])
		}
	}

	// Нормализация и вычисление энергии активации
	magnitude := cmplx.Abs(total) / float64(len(input)*len(dendrites))
	return magnitude * magnitude
}
