package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/aboutbrain/redozubov-model/config"
	"github.com/aboutbrain/redozubov-model/cortex"
	"github.com/aboutbrain/redozubov-model/learning"
	"github.com/aboutbrain/redozubov-model/utils"
)

// Добавляем функцию логирования состояния коры
func printCortexState(c *cortex.Cortex) {
	fmt.Println("\nСостояние коры:")
	for i, row := range c.Columns {
		for j, col := range row {
			fmt.Printf("  Колонка[%d][%d]: активирована=%t, уровень активации=%.4f\n",
				i, j, col.Activated, col.Activation)
		}
	}
	fmt.Printf("Глобальный контекст: %.4f\n", c.GlobalContext)
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("=== Инициализация модели коры ===")
	fmt.Printf("Размер паттерна: %d, нейронов в колонке: %d, порог активации: %.2f\n",
		config.PatternSize, config.NumNeurons, config.ActivationThreshold)

	cortex := cortex.NewCortex(3, 3)

	fmt.Println("\nГенерация входных данных...")
	input := utils.GenerateInputTensor(3, 3)

	fmt.Println("\n=== Первый проход обработки ===")
	cortex.ProcessInput(input)
	printCortexState(cortex)

	fmt.Println("\nПрименяем обучение...")
	applyLearning(cortex, input, cortex.GlobalContext)

	fmt.Println("\nПрименяем дофаминовую модуляцию...")
	for i, row := range cortex.Columns {
		for j, col := range row {
			if col.Activated {
				fmt.Printf("  Колонка[%d][%d] получает подкрепление\n", i, j)
				learning.ApplyDopamine(col, 1.5)
			}
		}
	}

	fmt.Println("\nКонсолидация памяти...")
	consolidateMemory(cortex)

	fmt.Println("\n=== Повторный проход после обучения ===")
	cortex.ProcessInput(input)
	printCortexState(cortex)

	// Добавляем детализацию по активированным колонкам
	fmt.Println("\nДетали активированных колонок:")
	for i, row := range cortex.Columns {
		for j, col := range row {
			if col.Activated {
				fmt.Printf("Колонка[%d][%d]:\n", i, j)
				for n, neuron := range col.Neurons {
					activation := neuron.CalculateActivation(input[i][j], cortex.GlobalContext)
					fmt.Printf("  Нейрон %d: активация=%.4f\n", n, activation)
				}
			}
		}
	}
}

func applyLearning(c *cortex.Cortex, input [][][]complex128, globalContext []float64) {
	for i, row := range c.Columns {
		for j, col := range row {
			if col.Activated && i < len(input) && j < len(input[i]) {
				for _, neuron := range col.Neurons {
					learning.ApplyHebbianLearning(neuron, input[i][j], globalContext)
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

	sum := 0.0
	for _, val := range c.GlobalContext {
		sum += val * val
	}
	sum = math.Sqrt(sum)
	if sum > 0 {
		for i := range c.GlobalContext {
			c.GlobalContext[i] /= sum
		}
	}
}
