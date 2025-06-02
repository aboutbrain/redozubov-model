package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/aboutbrain/redozubov-model/cortex"
	"github.com/aboutbrain/redozubov-model/learning"
	"github.com/aboutbrain/redozubov-model/utils"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	// Создаем кору 3x3
	cortex := cortex.NewCortex(3, 3)

	// Генерируем входные данные
	input := utils.GenerateInputTensor(3, 3)

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
