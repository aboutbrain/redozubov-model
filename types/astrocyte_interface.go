package types

// AstrocyteInterface определяет поведение астроцита
type AstrocyteInterface interface {
	TransferEnergy() float64
	GetCalciumLevel() float64
}
