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

func main() {
	rand.Seed(time.Now().UnixNano())

	cortex := cortex.NewCortex(3, 3)
	input := utils.GenerateInputTensor(3, 3)

	fmt.Println("=== Первый проход ===")
	cortex.ProcessInput(input)
	printCortexState(cortex)

	applyLearning(cortex, input, cortex.GlobalContext)

	fmt.Println("\nПрименяем дофаминовую модуляцию...")
	for _, row := range cortex.Columns {
		for _, col := range row {
			if col.Activated {
				learning.ApplyDopamine(col, 1.2)
			}
		}
	}

	fmt.Println("Консолидация памяти...")
	consolidateMemory(cortex)

	fmt.Println("\n=== После обучения ===")
	cortex.ProcessInput(input)
	printCortexState(cortex)
}

func printCortexState(c *cortex.Cortex) {
	for i, row := range c.Columns {
		for j, col := range row {
			fmt.Printf("Колонка[%d][%d]: активирована=%t, уровень=%.2f\n", i, j, col.Activated, col.Activation)
		}
	}
	fmt.Printf("Глобальный контекст: %.2f\n", c.GlobalContext)
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
