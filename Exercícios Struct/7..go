package main

import "fmt"

// Struct Animal
type Animal struct {
	nome    string
	especie string
	idade   int
	som     string
}

// Função para modificar o som que o animal faz
func modificarSom(a *Animal, novoSom string) {
	a.som = novoSom
}

// Função para imprimir as informações do animal e o som que ele faz
func imprimirInformacoesAnimal(a Animal) {
	fmt.Println("Nome:", a.nome)
	fmt.Println("Espécie:", a.especie)
	fmt.Println("Idade:", a.idade)
	fmt.Println("Som:", a.som)
}

func main() {
	animal := Animal{
		nome:    "Rex",
		especie: "Cachorro",
		idade:   5,
		som:     "Latido",
	}

	imprimirInformacoesAnimal(animal)
	fmt.Println()

	modificarSom(&animal, "Miado")
	imprimirInformacoesAnimal(animal)
}
