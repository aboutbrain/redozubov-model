package cortex

import (
	"github.com/aboutbrain/redozubov-model/astrocyte"
	"github.com/aboutbrain/redozubov-model/minicolumn"
	"github.com/aboutbrain/redozubov-model/types"
)

type Cortex struct {
	Columns       [][]*minicolumn.Minicolumn
	GlobalContext []float64
	Astrocytes    []*astrocyte.Astrocyte
}

func NewCortex(rows, cols int) *Cortex {
	c := &Cortex{
		Columns:       make([][]*minicolumn.Minicolumn, rows),
		GlobalContext: make([]float64, minicolumn.PatternSize),
	}

	totalColumns := rows * cols
	numAstrocytes := totalColumns / 4
	if numAstrocytes < 1 {
		numAstrocytes = 1
	}
	c.Astrocytes = make([]*astrocyte.Astrocyte, numAstrocytes)
	for i := range c.Astrocytes {
		c.Astrocytes[i] = astrocyte.NewAstrocyte()
	}

	astroIndex := 0
	for i := range c.Columns {
		c.Columns[i] = make([]*minicolumn.Minicolumn, cols)
		for j := range c.Columns[i] {
			currentAstrocyte := c.Astrocytes[astroIndex]
			c.Columns[i][j] = minicolumn.NewMinicolumn(types.AstrocyteInterface(currentAstrocyte))

			if (i*cols+j+1)%4 == 0 && astroIndex < len(c.Astrocytes)-1 {
				astroIndex++
			}
		}
	}

	return c
}

func (c *Cortex) UpdateAstrocytes() {
	for _, astro := range c.Astrocytes {
		astro.Update()
	}
}

func (c *Cortex) ProcessInput(input [][][]complex128) {
	for i, row := range c.Columns {
		for j, col := range row {
			if i < len(input) && j < len(input[i]) {
				// Убрали создание fullContext - теперь используем только GlobalContext
				col.ProcessPattern(input[i][j], c.GlobalContext)

				if col.Activated {
					c.GlobalContext[0] = 0.9
				}
			}
		}
	}

	// Добавляем конкурентное торможение
	LateralInhibition(c.Columns)
}

func (c *Cortex) ApplyAttention(focus [][]float64) {
	for i, row := range c.Columns {
		for j, col := range row {
			if i < len(focus) && j < len(focus[i]) {
				// Усиливаем активацию в фокусе внимания
				col.Activation *= 1.0 + focus[i][j]
			}
		}
	}
}

func (c *Cortex) RestCycle() {
	for _, row := range c.Columns {
		for _, col := range row {
			col.Rest()
		}
	}
}
