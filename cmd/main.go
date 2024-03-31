package main

import (
	"geneticsAI"
	"time"
)

func main() {

	//t := time.Now().UnixMilli()
	generation := geneticsAI.InitGeneration(200)
	//for true {
	generation.Mutation()
	generation.All = geneticsAI.SortFitness(generation.All)
	//}
	println(len(generation.All))
	time.Sleep(400 * time.Millisecond)
}
