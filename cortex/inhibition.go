package cortex

import (
	"github.com/aboutbrain/redozubov-model/config"
	"github.com/aboutbrain/redozubov-model/minicolumn"
)

func (c *Cortex) LateralInhibition() {
	maxActivation := 0.0
	var maxCol *minicolumn.Minicolumn

	// Находим наиболее активную колонку
	for _, row := range c.Columns {
		for _, col := range row {
			if col.Activation > maxActivation {
				maxActivation = col.Activation
				maxCol = col
			}
		}
	}

	// Применяем торможение
	if maxCol != nil && maxCol.Activation > 0.5 {
		for _, row := range c.Columns {
			for _, col := range row {
				if col != maxCol {
					col.Activation *= 0.2 // Ослабляем соседей
					if col.Activation < config.ActivationThreshold {
						col.Activated = false
					}
				}
			}
		}
	}
}
