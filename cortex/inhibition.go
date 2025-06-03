package cortex

import "github.com/aboutbrain/redozubov-model/minicolumn"

func (c *Cortex) LateralInhibition() {
	for _, row := range c.Columns {
		maxActivation := 0.0
		// Находим максимальную активацию
		for _, col := range row {
			if col.Activation > maxActivation {
				maxActivation = col.Activation
			}
		}
		// Подавляем слабые колонки
		for _, col := range row {
			if col.Activation < maxActivation*0.7 {
				col.Activated = false
			}
		}
	}
}

func LateralInhibition(columns [][]*minicolumn.Minicolumn) {
	for i, row := range columns {
		for j, col := range row {
			if col.Activated {
				// Подавление соседей
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						if di == 0 && dj == 0 {
							continue
						}
						ni, nj := i+di, j+dj
						if ni >= 0 && ni < len(columns) && nj >= 0 && nj < len(row) {
							neighbor := columns[ni][nj]
							// Умеренное подавление
							neighbor.EnergyLevel *= 0.85
							neighbor.Activation *= 0.9
						}
					}
				}
			}
		}
	}
}
