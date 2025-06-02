package utils

import (
	"math/rand"

	"github.com/aboutbrain/redozubov-model/config"
)

func GeneratePattern() []complex128 {
	pattern := make([]complex128, config.PatternSize)
	for i := range pattern {
		pattern[i] = complex(rand.NormFloat64(), rand.NormFloat64())
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
