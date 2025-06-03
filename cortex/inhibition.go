package cortex

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
