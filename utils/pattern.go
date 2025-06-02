package utils

import (
	"github.com/aboutbrain/redozubov-model/minicolumn"
	"math/rand"
)

func GeneratePattern() []complex128 {
	pattern := make([]complex128, minicolumn.PatternSize)
	for i := range pattern {
		// Увеличиваем амплитуду в 10 раз
		real := rand.NormFloat64() * 10
		imag := rand.NormFloat64() * 10
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
