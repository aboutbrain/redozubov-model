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
	for i := range columns {
		for j := range columns[i] {
			if columns[i][j].Activated {
				// Подавляем соседние колонки
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						ni, nj := i+di, j+dj
						if ni >= 0 && ni < len(columns) &&
							nj >= 0 && nj < len(columns[0]) &&
							!(di == 0 && dj == 0) {
							columns[ni][nj].EnergyLevel *= 0.7
						}
					}
				}
			}
		}
	}
}
