package cortex

import (
	"fmt"
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
			fmt.Printf("Создана миниколонка [%d][%d]\n", i, j) // Добавляем логирование
		}
	}

	// Инициализация глобального контекста
	for i := range c.GlobalContext {
		c.GlobalContext[i] = rand.NormFloat64() * 0.2
	}

	fmt.Println("Кора инициализирована успешно")
	return c
}

func (c *Cortex) ProcessInput(input [][][]complex128) {
	fmt.Println("Начало обработки входных данных...")
	for i, row := range c.Columns {
		for j, col := range row {
			if i < len(input) && j < len(input[i]) {
				fmt.Printf("Обработка колонки [%d][%d]\n", i, j)
				col.ProcessPattern(input[i][j], c.GlobalContext)
			}
		}
	}

	fmt.Println("Применение латерального торможения...")
	c.LateralInhibition()
}
