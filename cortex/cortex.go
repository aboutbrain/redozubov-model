package cortex

import (
	"math/rand"

	"github.com/aboutbrain/redozubov-model/config"
	"github.com/aboutbrain/redozubov-model/minicolumn"
)

type Cortex struct {
	Columns       [][]*minicolumn.Minicolumn
	GlobalContext []float64
}

func NewCortex(rows, cols int) *Cortex {
	c := &Cortex{
		Columns:       make([][]*minicolumn.Minicolumn, rows),
		GlobalContext: make([]float64, config.PatternSize),
	}

	for i := range c.Columns {
		c.Columns[i] = make([]*minicolumn.Minicolumn, cols)
		for j := range c.Columns[i] {
			c.Columns[i][j] = minicolumn.NewMinicolumn()
		}
	}

	// Инициализация глобального контекста
	for i := range c.GlobalContext {
		c.GlobalContext[i] = rand.NormFloat64() * 0.2
	}

	return c
}

func (c *Cortex) ProcessInput(input [][][]complex128) {
	for i, row := range c.Columns {
		for j, col := range row {
			if i < len(input) && j < len(input[i]) {
				col.ProcessPattern(input[i][j], c.GlobalContext)
			}
		}
	}

	c.LateralInhibition()
}
