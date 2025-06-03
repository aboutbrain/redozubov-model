package utils

import (
	"github.com/aboutbrain/redozubov-model/minicolumn"
	"math"
	"math/rand"
)

func GeneratePattern() []complex128 {
	pattern := make([]complex128, minicolumn.PatternSize)
	for i := range pattern {
		// Более структурированные паттерны
		phase := float64(i) * 2 * math.Pi / float64(minicolumn.PatternSize)
		real := math.Cos(phase) + rand.NormFloat64()*0.1
		imag := math.Sin(phase) + rand.NormFloat64()*0.1
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
