package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aboutbrain/redozubov-model/cortex"
	"github.com/aboutbrain/redozubov-model/learning"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Создаем кору 3x3
	cortex := cortex.NewCortex(3, 3)

	// Фиксированные тестовые данные вместо случайных
	input := [][][]complex128{
		{
			{complex(5, 2), complex(3, 4), complex(1, 6), complex(7, 1), complex(2, 5)},
			{complex(4, 3), complex(6, 2), complex(2, 7), complex(5, 4), complex(3, 6)},
			{complex(3, 5), complex(2, 6), complex(4, 3), complex(6, 2), complex(5, 4)},
		},
		{
			{complex(6, 1), complex(2, 7), complex(5, 3), complex(3, 5), complex(4, 4)},
			{complex(1, 8), complex(7, 2), complex(3, 6), complex(4, 5), complex(6, 3)},
			{complex(2, 7), complex(5, 4), complex(6, 3), complex(1, 8), complex(7, 2)},
		},
		{
			{complex(7, 3), complex(4, 5), complex(2, 8), complex(6, 4), complex(3, 7)},
			{complex(3, 6), complex(1, 9), complex(7, 3), complex(2, 8), complex(5, 5)},
			{complex(4, 7), complex(6, 4), complex(3, 7), complex(5, 5), complex(1, 9)},
		},
	}

	for step := 0; step < 5; step++ {
		fmt.Printf("\n=== Шаг %d ===\n", step)

		// Обновляем астроциты
		cortex.UpdateAstrocytes()

		// Обрабатываем вход
		cortex.ProcessInput(input)

		// Выводим состояние коры
		printCortexState(cortex)

		// Применяем обучение
		applyLearning(cortex, input)

		// Применяем дофаминовую модуляцию
		fmt.Println("\nПрименяем дофаминовую модуляцию...")
		for _, row := range cortex.Columns {
			for _, col := range row {
				if col.Activated {
					learning.ApplyDopamine(col, 1.2)
				}
			}
		}

		// Консолидация памяти
		fmt.Println("Консолидация памяти...")
		consolidateMemory(cortex)
	}
}

func printCortexState(c *cortex.Cortex) {
	for i, row := range c.Columns {
		for j, col := range row {
			fmt.Printf("Колонка[%d][%d]: активирована=%t, уровень=%.2f, энергия=%.2f\n",
				i, j, col.Activated, col.Activation, col.EnergyLevel)
		}
	}
}

func applyLearning(c *cortex.Cortex, input [][][]complex128) {
	for i, row := range c.Columns {
		for j, col := range row {
			if col.Activated && i < len(input) && j < len(input[i]) {
				// Убрали создание fullContext - теперь используем только GlobalContext
				for _, neuron := range col.Neurons {
					learning.ApplyHebbianLearning(neuron, input[i][j], c.GlobalContext)
				}
			}
		}
	}
}

func consolidateMemory(c *cortex.Cortex) {
	for _, row := range c.Columns {
		for _, col := range row {
			for _, neuron := range col.Neurons {
				learning.Consolidate(neuron)
			}
		}
	}
}
