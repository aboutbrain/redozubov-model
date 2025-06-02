package utils

import (
	"github.com/aboutbrain/redozubov-model/config"
	"math"
	"math/rand"
)

func GeneratePattern() []complex128 {
	pattern := make([]complex128, config.PatternSize)
	for i := range pattern {
		real := rand.NormFloat64()
		imag := rand.NormFloat64()
		// Нормализуем, чтобы не было слишком больших значений
		magnitude := math.Sqrt(real*real + imag*imag)
		if magnitude > 1e-5 {
			real /= magnitude
			imag /= magnitude
		}
		pattern[i] = complex(real, imag)
	}
	return pattern
}

func GenerateInputTensor(rows, cols int) [][][]complex128 {
	input := make([][][]complex128, rows)
	for i := range input {
		input[i] = make([][]complex128, cols)
		for j := range input[i] {
			input[i][j] = GeneratePattern()
		}
	}
	return input
}
