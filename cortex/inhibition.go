package cortex

import (
	"github.com/aboutbrain/redozubov-model/config"
	"math"
)

// LateralInhibition применяет конкурентное торможение в слое коры
func (c *Cortex) LateralInhibition() {
	for i, row := range c.Columns {
		for j, col := range row {
			if col.Activation > 0.7 {
				// Подавление соседних колонок
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						ni, nj := i+di, j+dj
						if ni >= 0 && ni < len(c.Columns) &&
							nj >= 0 && nj < len(row) &&
							!(di == 0 && dj == 0) {

							// Уменьшаем активацию соседей
							neighbor := c.Columns[ni][nj]
							neighbor.Activation *= math.Exp(-float64(di*di + dj*dj))
							if neighbor.Activation < config.ActivationThreshold {
								neighbor.Activated = false
							}
						}
					}
				}
			}
		}
	}
}
