package main

import (
	"fmt"
	"geneticsAI"
	"time"
)

func main() {

	//t := time.Now().UnixMilli()
	generation := geneticsAI.InitGeneration(200)
	first := generation.All[0].CountPeople
	for i := 0; true; i++ {
		if generation.All[0].CountPeople != first {
			first = generation.All[0].CountPeople
			println("[LOG]", time.Now().Format("01-02-2006 15:04:05"), "|", first, "|", "  new   ", "|", "gen:", fmt.Sprintf("%dº", i))
		}
		generation.Mutation()
		generation.All = geneticsAI.SortFitness(generation.All)
		generation.Crossover()
		generation.All = generation.All[:200]
		if generation.All[0].CountPeople-generation.All[199].CountPeople == 0 {
			generation.All = append(generation.All[:3], geneticsAI.InitGeneration(200).All[3:]...)
			println("[LOG]", time.Now().Format("01-02-2006 15:04:05"), "|", first, "|", "disaster", "|", "gen:", fmt.Sprintf("%dº", i))
		}
	}
}
