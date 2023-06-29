package main

//Crie uma struct chamada Retângulo com os cmapos "largura"e "altura"" Escreva uma função que receba um Retângulo
//como parâmetro e calcule a área do retângulo

import "fmt"

type Retângulo struct {
	base   float64
	altura float64
}

func ÁreaRetângulo(x Retângulo) float64 {
	return x.base * x.altura
}

func main() {

	var retângulo_aleatório Retângulo

	fmt.Println("Qual a base do retângulo?")
	fmt.Scan(&retângulo_aleatório.base)
	fmt.Println("Qual a altura do retângulo")
	fmt.Scan(&retângulo_aleatório.altura)

	ÁreaRetângulo()

}
