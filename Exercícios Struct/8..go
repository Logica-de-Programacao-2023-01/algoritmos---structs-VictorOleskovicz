package main

import "fmt"

// Struct Viagem
type Viagem struct {
	origem  string
	destino string
	data    string
	preco   float64
}

// Função para encontrar a viagem mais cara
func encontrarViagemMaisCara(viagens []Viagem) Viagem {
	viagemMaisCara := viagens[0]

	for _, viagem := range viagens {
		if viagem.preco > viagemMaisCara.preco {
			viagemMaisCara = viagem
		}
	}

	return viagemMaisCara
}

func main() {
	viagens := []Viagem{
		{origem: "São Paulo", destino: "Rio de Janeiro", data: "2023-07-01", preco: 200.0},

