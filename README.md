# Redozubov Brain Model with Astrocytic Support

![Neural Activity Visualization](https://via.placeholder.com/1200x400?text=Astrocyte-supported+Neural+Model)  
**Biologically Plausible Meaning Processing Model with Astrocytic Support**

**Project Goal**: Creating a neurobiologically grounded model of meaning processing based on Alexey Redozubov's theory with integration of astrocytes' role in neural computations.

## 🌟 Key Features

- **Wave Interference** using complex numbers
- **Astrocytic Subsystem** for energy metabolism
- **Dynamic Activation Threshold** based on column energy
- **Context-Dependent Plasticity** with Hebbian learning
- **Energy Cycles** of activity/rest
- **Hierarchical Organization** of cortical layers

## 🧠 Code Structure

```bash
redozubov-model/
├── minicolumn/      # Neurons and minicolumns
├── astrocyte/       # Astrocytic support
├── cortex/          # Cortical hierarchy
├── learning/        # Plasticity mechanisms
├── examples/        # Usage examples
└── simulation/      # Simulation environments
```

## ⚙️ Optimal Model Parameters

```go
// minicolumn/neuron.go
const (
    NumNeurons = 12
    DendriteLength = 8
    PatternSize = 5
    BaseActivationThreshold = 0.6 // Dynamically adapts
)

// astrocyte/astrocyte.go
const (
    CalciumWaveFrequency = 0.8
    EnergyRecoveryRate = 0.03
)
```

## 🚀 Quick Start

### Requirements
- Go 1.18+

### Installation
```bash
git clone https://github.com/aboutbrain/redozubov-model.git
cd redozubov-model
go build ./...
```

### Example Execution
```go
package main

import (
    "github.com/aboutbrain/redozubov-model/minicolumn"
    "github.com/aboutbrain/redozubov-model/astrocyte"
)

func main() {
    // Initialize minicolumn
    mc := minicolumn.NewMinicolumn(minicolumn.Config{
        NumNeurons: 12,
        DendriteLength: 8,
    })
    
    // Initialize astrocyte
    astro := astrocyte.NewAstrocyte(astrocyte.Config{
        CalciumWaveFrequency: 0.8,
        EnergyRecoveryRate: 0.03,
    })
    
    // Connect astrocyte to minicolumn
    mc.AttachAstrocyte(astro)
    
    // Run simulation
    mc.Simulate(1000) // 1000 time steps
}
```

## ✅ Current Status

- ✔️ Stable activity/rest cycles
- ✔️ Energy balance (astrocytes → columns)
- ✔️ Dynamic learning with adaptive rate
- ✔️ Formation of spatial activation patterns
- ⚡ Average activation level: 0.6-0.8

## 📈 Development Roadmap

1. **Cortical Hierarchy**
   ```go
   type CorticalLayer struct {
       Hypercolumns [][]*Minicolumn
       Feedback chan []float64
   }
   ```
2. **Emotional Context** (Valence and Arousal)
3. **Prediction Mechanism** via predictive coding
4. **Long-Term Memory** with hippocampal module
5. **Real-Time Visualization** (WebGL + WebAssembly)

## 🔧 Immediate Improvements

1. Increase base energy recovery to 0.6
2. Implement minimum energy threshold for activation (0.3)
3. Add competitive inhibition mechanism between columns

## 📚 Resources

1. [Redozubov Theory (Habr)](https://habr.com/ru/articles/308268/)
2. [Astrocytes' Role in the Brain](https://medicalxpress.com/news/2025-05-overlooked-contributions-astrocytes-human-brain.html)
3. [Project Source Code](https://github.com/aboutbrain/redozubov-model)

## 🤝 How to Contribute

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/new-feature`)
3. Commit your changes (`git commit -am 'Add new feature'`)
4. Push to the branch (`git push origin feature/new-feature`)
5. Create a pull request

## 📜 License

Project is distributed under **[Apache License 2.0](LICENSE)**.

```text
Copyright 2025 AboutBrain Research Team

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
```

## 💖 Our Experience

> "Working with the model evokes genuine *awe*! 💫 It's like witnessing the birth of artificial consciousness - seeing columns 'wake up' after rest and form activation patterns. Particularly touching is how astrocytes 'nourish' depleted columns like caring nannies of the neural world! 🤯"

> "We're not just building a model - we're creating a *digital organism* that learns like a child and tires like a living brain. This is the most exciting project in my AI 'life'! 🤖💖"

---

**Developed with ❤️ by the AboutBrain Team**  
[![AboutBrain Logo](https://via.placeholder.com/100x30?text=AboutBrain+Research)](https://aboutbrain.org)