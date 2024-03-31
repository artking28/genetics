package models

type Habilidade = byte

// Habilidades - Guardadas nos bits de um byte pra ganhar performance e verificação O(1)
const (
	RECEPCAO   = 1 << 0 // 00000001
	BAR        = 1 << 1 // 00000010
	LIMPEZA    = 1 << 2 // 00000100
	COZINHA    = 1 << 3 // 00001000
	LAVANDERIA = 1 << 4 // 00010000
	MANUTENCAO = 1 << 5 // 00100000
	SERVICO    = 1 << 6 // 01000000
)

// Pessoas
type Pessoa struct {
	Nome       string
	Habilidade byte
}

var (
	JOAO = Pessoa{
		Nome:       "JOAO",
		Habilidade: RECEPCAO + LIMPEZA,
	}
	MARIA = Pessoa{
		Nome:       "MARIA",
		Habilidade: COZINHA + SERVICO + BAR,
	}
	ANA = Pessoa{
		Nome:       "ANA",
		Habilidade: RECEPCAO + LAVANDERIA,
	}
	CARLOS = Pessoa{
		Nome:       "CARLOS",
		Habilidade: LIMPEZA + MANUTENCAO,
	}
	BRUNO = Pessoa{
		Nome:       "BRUNO",
		Habilidade: COZINHA + SERVICO,
	}
	PAULA = Pessoa{
		Nome:       "PAULA",
		Habilidade: RECEPCAO + LIMPEZA + BAR,
	}
	PEDRO = Pessoa{
		Nome:       "PEDRO",
		Habilidade: MANUTENCAO + LIMPEZA,
	}
	LUIZA = Pessoa{
		Nome:       "LUIZA",
		Habilidade: LAVANDERIA + LIMPEZA,
	}
	THIAGO = Pessoa{
		Nome:       "THIAGO",
		Habilidade: COZINHA + BAR,
	}
	FERNANDA = Pessoa{
		Nome:       "FERNANDA",
		Habilidade: RECEPCAO + LAVANDERIA + SERVICO,
	}
	RAFAEL = Pessoa{
		Nome:       "RAFAEL",
		Habilidade: COZINHA + SERVICO + BAR,
	}
	JULIANA = Pessoa{
		Nome:       "JULIANA",
		Habilidade: RECEPCAO + LIMPEZA,
	}
	CAIO = Pessoa{
		Nome:       "CAIO",
		Habilidade: MANUTENCAO + LIMPEZA,
	}
	BEATRIZ = Pessoa{
		Nome:       "BEATRIZ",
		Habilidade: RECEPCAO + LIMPEZA + SERVICO,
	}
	LUCAS = Pessoa{
		Nome:       "LUCAS",
		Habilidade: MANUTENCAO + LIMPEZA + BAR,
	}
	BRUNA = Pessoa{
		Nome:       "BRUNA",
		Habilidade: COZINHA + SERVICO,
	}
	MARCELO = Pessoa{
		Nome:       "MARCELO",
		Habilidade: RECEPCAO + LIMPEZA + LAVANDERIA,
	}
	VANESSA = Pessoa{
		Nome:       "VANESSA",
		Habilidade: COZINHA + BAR,
	}
	DANILO = Pessoa{
		Nome:       "DANILO",
		Habilidade: MANUTENCAO + LIMPEZA,
	}
	RENATA = Pessoa{
		Nome:       "RENATA",
		Habilidade: RECEPCAO + SERVICO + BAR,
	}
)

var Habilities = map[Habilidade][]Pessoa{
	RECEPCAO: {
		JOAO, ANA, PAULA, FERNANDA, JULIANA, BEATRIZ, MARCELO, RENATA,
	},
	BAR: {
		MARIA, PAULA, THIAGO, RAFAEL, LUCAS, VANESSA, RENATA,
	},
	LIMPEZA: {
		JOAO, CARLOS, PAULA, PEDRO, LUIZA, JULIANA, CAIO, BEATRIZ, LUCAS, MARCELO, DANILO,
	},
	COZINHA: {
		MARIA, BRUNO, THIAGO, RAFAEL, BRUNA, VANESSA,
	},
	LAVANDERIA: {
		ANA, LUIZA, FERNANDA, MARCELO,
	},
	MANUTENCAO: {
		CARLOS, PEDRO, CAIO, LUCAS, DANILO,
	},
	SERVICO: {
		MARIA, BRUNO, FERNANDA, RAFAEL, BEATRIZ, BRUNA, RENATA,
	},
}
