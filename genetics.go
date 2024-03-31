package geneticsAI

import (
	"geneticsAI/models"
	"math/rand"
)

type Generation struct {
	All []models.Individuo
	Gen int
}

func InitGeneration(size int) Generation {
	var ret Generation
	for i := 0; i < size; i++ {
		k := models.InitIndividuo()
		//if k.CountPeople > cMax {
		//	cMax = k.CountPeople
		//}
		//if k.CountPeople < cMin {
		//	cMin = k.CountPeople
		//}
		//c += k.CountPeople
		//println(i)
		ret.All = append(ret.All, k)
	}
	ret.Gen = 1
	return ret
}

func (this *Generation) Mutation() {

	novos, parciais := 10, 30

	// Cria dez novos individuo
	for l := 0; l < novos; l++ {
		this.All = append(this.All, models.InitIndividuo())
	}

	// Modifica 30 indivíduos parcialmente
	for l := 0; l < parciais; l++ {

		parcial := this.All[rand.Int()%len(this.All)]
		outter := 0
		lastTurnoPeople0, lastTurnoPeople1 := map[models.Pessoa]int{}, map[models.Pessoa]int{}
		for i := 69; i < 90; {
			t := models.InitTurno()

			out, commom := false, 0
			//Ve se a pessoa n esta trabalhando
			for k := range t.People {
				if lastTurnoPeople0[k] == 1 {
					out = true
					if outter >= 3 {
						if i >= 2 {
							lastTurnoPeople0 = parcial.Content[i-2].People
							lastTurnoPeople1 = parcial.Content[i-1].People
							i--
						} else {
							lastTurnoPeople0, lastTurnoPeople1 = map[models.Pessoa]int{}, map[models.Pessoa]int{}
							i = 0
						}
						outter, commom = 0, 0
						continue
					}
					outter++
					continue
				} else if lastTurnoPeople0[k] == 0 && lastTurnoPeople1[k] == 1 {
					commom++
				}
			}
			if out {
				continue
			}
			parcial.Content[i] = t
			parcial.CountPeople += len(t.People) - commom
			lastTurnoPeople0 = lastTurnoPeople1
			lastTurnoPeople1 = t.People
			i++
		}
		this.All = append(this.All, parcial)
	}
}

func (this *Generation) Crossover() {

}

func SortFitness(inds []models.Individuo) (ret []models.Individuo) {
	if len(inds) < 2 {
		return inds
	}
	pivot := inds[rand.Int()%len(inds)]
	var smaller, greater []models.Individuo
	for i := 0; i < len(inds)-1; i++ {
		if inds[i].Calc() > inds[i+1].Calc() {
			greater = append(greater, inds[i])
			if i+2 == len(inds) {
				smaller = append(smaller, inds[i+1])
			}
			continue
		}
		smaller = append(smaller, inds[i])
		if i+2 == len(inds) {
			greater = append(greater, inds[i+1])
		}
	}
	ret = append(ret, SortFitness(smaller)...)
	ret = append(ret, pivot)
	return append(ret, SortFitness(greater)...)
}
