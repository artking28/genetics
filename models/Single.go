package models

import (
	"math/rand"
	"strings"
)

// IndSize Quantidade de genes em cada indivíduo
const IndSize = 90

type (
	Individuo struct {
		Content     [90]Turno
		CountPeople int
	}

	Turno struct {
		People map[Pessoa]int
		Novas  int
	}
)

// InitIndividuo - Inicia um individuo com preenchimento aleatório e mantendo as regras
func InitIndividuo() Individuo {
	var ret Individuo
	ret.Content = [IndSize]Turno{}
	outter := 0
	lastTurnoPeople0, lastTurnoPeople1 := map[Pessoa]int{}, map[Pessoa]int{}
	for i := 0; i < IndSize; {
		t := InitTurno()

		out, commom := false, 0
		//Ve se a pessoa n esta trabalhando
		for k := range t.People {
			if lastTurnoPeople0[k] == 1 {
				out = true
				if outter >= 3 {
					if i >= 2 {
						lastTurnoPeople0 = ret.Content[i-2].People
						lastTurnoPeople1 = ret.Content[i-1].People
						i--
					} else {
						lastTurnoPeople0, lastTurnoPeople1 = map[Pessoa]int{}, map[Pessoa]int{}
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
		t.Novas = len(t.People) - commom
		ret.Content[i] = t
		lastTurnoPeople0 = lastTurnoPeople1
		lastTurnoPeople1 = t.People
		i++
	}
	for _, i := range ret.Content {
		ret.CountPeople += i.Novas
	}
	return ret
}

func (this *Individuo) Calc() float64 {
	return float64((1/this.CountPeople)*10) - float64(this.CountPeople)
}

func (this *Individuo) UpdateCount() {
	this.CountPeople = 0
	for i := 0; i < len(this.Content); i++ {
		commom := 0
		if i == 0 {
			this.CountPeople += len(this.Content[i].People)
			continue
		}
		for key := range this.Content[i].People {
			if this.Content[i-1].People[key] == 1 {
				commom++
			}
		}
		this.CountPeople += len(this.Content[i].People) - commom
	}
}

// InitTurno - Inicia um turno com preenchimento aleatório
func InitTurno() (ret Turno) {

	// Todos os bits de cada uma das 7 habilidades
	habilidadesFaltando := byte(0b01111111)
	ret.People = map[Pessoa]int{}
	// Itero as habilidades
	for i := 0; i < 7; i++ {
		habilidade := byte(1 << i)
		// Habilidade ja resolvida
		if habilidadesFaltando&habilidade != habilidade {
			continue
		}
		// Pego as pessoas de certa habilidade
		capazes := Habilities[habilidade]
		// Pego uma pessoa aleatória e q ja n esteja na lista
		pessoa := capazes[rand.Int()%len(capazes)]
		for ret.People[pessoa] == 1 {
			pessoa = capazes[rand.Int()%len(capazes)]
		}
		// Faco o xor com as habilidades da pessoa pra retirar das pendências
		habilidadesFaltando ^= pessoa.Habilidade
		ret.People[pessoa] = 1
	}
	return ret
}

// Println - Só para printar turno na tela
func (this *Turno) Println() {
	str := strings.Builder{}
	str.WriteString("[Turno] ~> {")
	for k := range this.People {
		str.WriteString(", " + k.Nome)
	}
	str.WriteByte(' ')
	str.WriteByte('}')
	println(strings.Replace(str.String(), ",", "", 1))
}
