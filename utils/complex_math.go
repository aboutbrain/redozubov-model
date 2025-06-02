package utils

import (
	"math"
	"math/cmplx"
)

// ComplexNorm нормализует комплексный вектор
func ComplexNorm(vec []complex128) {
	mag := 0.0
	for _, v := range vec {
		mag += cmplx.Abs(v) * cmplx.Abs(v)
	}
	mag = math.Sqrt(mag)

	if mag > 0 {
		for i := range vec {
			vec[i] /= complex(mag, 0)
		}
	}
}

// ComplexDotProduct вычисляет скалярное произведение комплексных векторов
func ComplexDotProduct(a, b []complex128) complex128 {
	result := 0i
	for i := range a {
		result += a[i] * cmplx.Conj(b[i])
	}
	return result
}
