package utils

import (
	"math"
)

func ComplexNorm(vec []complex128) {
	mag := 0.0
	for _, v := range vec {
		mag += real(v)*real(v) + imag(v)*imag(v)
	}
	mag = math.Sqrt(mag)

	if mag > 0 {
		for i := range vec {
			vec[i] = complex(real(vec[i])/mag, imag(vec[i])/mag)
		}
	}
}
