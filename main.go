package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/aboutbrain/redozubov-model/cortex"
	"github.com/aboutbrain/redozubov-model/learning"
	"github.com/aboutbrain/redozubov-model/utils"
)

const (
	InitialLearningRate = 0.3
	MinLearningRate     = 0.1
)

// Динамическая регулировка скорости обучения
func adjustLearningRate(step int) float64 {
	// Экспоненциальное затухание скорости обучения
	return math.Max(MinLearningRate, InitialLearningRate*math.Exp(-float64(step)/20))
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// Создаем кору 3x3
	cortex := cortex.NewCortex(3, 3)

	// Генерируем входные данные
	input := utils.GenerateInputTensor(3, 3)

	// Статистика
	totalActivations := 0
	maxActivation := 0.0

	for step := 0; step < 20; step++ {
		fmt.Printf("\n=== Шаг %d ===\n", step)

		// Обновляем астроциты
		cortex.UpdateAstrocytes()

		// Обрабатываем вход
		cortex.ProcessInput(input)

		// Выводим состояние коры
		printCortexState(cortex)

		// Применяем обучение с динамической скоростью
		currentLR := adjustLearningRate(step)
		applyLearning(cortex, input, currentLR)

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

		// Сбор статистики
		stepActivations := 0
		for _, row := range cortex.Columns {
			for _, col := range row {
				if col.Activated {
					stepActivations++
					if col.Activation > maxActivation {
						maxActivation = col.Activation
					}
				}
			}
		}
		totalActivations += stepActivations

		fmt.Printf("Шаг %d: активировано %d/%d колонок (макс. активация: %.2f, LR: %.3f)\n",
			step, stepActivations, len(cortex.Columns)*len(cortex.Columns[0]), maxActivation, currentLR)

		// Визуализация энергии
		printEnergyGrid(cortex)

		// Каждые 5 шагов добавляем цикл отдыха
		if step%5 == 4 {
			fmt.Println("\n=== Цикл отдыха ===")
			cortex.RestCycle()
			cortex.UpdateAstrocytes()
			printEnergyGrid(cortex)
		}
	}

	fmt.Printf("\nИтого: активаций=%d, средняя=%.1f на шаг\n",
		totalActivations, float64(totalActivations)/20)

	// Уменьшим количество шагов для стабильности
	totalSteps := 15

	// Добавим разнообразие входных данных
	var inputs [][][][]complex128
	for i := 0; i < totalSteps; i++ {
		inputs = append(inputs, utils.GenerateInputTensor(3, 3))
	}

	for step := 0; step < totalSteps; step++ {
		// Используем разные входные данные на каждом шаге
		cortex.ProcessInput(inputs[step])

		// ... остальной код ...

		// Более частые циклы отдыха
		if step%3 == 2 {
			fmt.Println("\n=== Цикл отдыха ===")
			cortex.RestCycle()
			cortex.UpdateAstrocytes()
			printEnergyGrid(cortex)
		}
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

func printEnergyGrid(c *cortex.Cortex) {
	fmt.Println("Карта энергии (x10):")
	for _, row := range c.Columns {
		for _, col := range row {
			energyLevel := int(col.EnergyLevel * 10)
			fmt.Printf("%d ", energyLevel)
		}
		fmt.Println()
	}
}

func applyLearning(c *cortex.Cortex, input [][][]complex128, learningRate float64) {
	for i, row := range c.Columns {
		for j, col := range row {
			if col.Activated && i < len(input) && j < len(input[i]) {
				for _, neuron := range col.Neurons {
					learning.ApplyHebbianLearning(neuron, input[i][j], c.GlobalContext, learningRate)
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
