package astrocyte

import (
	"math"
	"math/rand"
	"time"

	"github.com/aboutbrain/redozubov-model/types"
)

const (
	CalciumWaveFrequency = 0.5 // Hz
	EnergyTransferRate   = 0.8
)

type Astrocyte struct {
	CalciumLevel  float64
	EnergyReserve float64
	LastActivity  time.Time
}

func NewAstrocyte() *Astrocyte {
	return &Astrocyte{
		CalciumLevel:  0.3,
		EnergyReserve: 1.0,
		LastActivity:  time.Now(),
	}
}

func (a *Astrocyte) Update() {
	elapsed := time.Since(a.LastActivity).Seconds()
	phase := 2 * math.Pi * CalciumWaveFrequency * elapsed
	// Увеличиваем амплитуду кальциевых волн
	a.CalciumLevel = 0.8 + 0.5*math.Sin(phase)

	a.CalciumLevel += rand.Float64()*0.1 - 0.05
	a.CalciumLevel = math.Max(0.0, math.Min(1.0, a.CalciumLevel))

	a.EnergyReserve = math.Min(1.0, a.EnergyReserve+0.01)
	a.LastActivity = time.Now()
}

func (a *Astrocyte) TransferEnergy() float64 {
	// Увеличиваем передачу энергии
	transferred := a.EnergyReserve * 0.9
	a.EnergyReserve -= transferred
	return transferred
}

func (a *Astrocyte) GetCalciumLevel() float64 {
	return a.CalciumLevel
}

// Проверка реализации интерфейса
var _ types.AstrocyteInterface = (*Astrocyte)(nil)
