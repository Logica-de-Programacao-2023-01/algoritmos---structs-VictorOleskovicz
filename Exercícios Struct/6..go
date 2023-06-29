package main

import "fmt"

// Struct Funcionário
type Funcionario struct {
	nome    string
	salario float64
	idade   int
}

// Função para aumentar o salário do funcionário em uma determinada porcentagem
func aumentarSalario(f *Funcionario, percentual float64) {
	f.salario *= (1 + percentual/100)
}

// Função para diminuir o salário do funcionário em uma determinada porcentagem
func diminuirSalario(f *Funcionario, percentual float64) {
	f.salario *= (1 - percentual/100)
}

// Função para calcular o tempo de serviço do funcionário
func calcularTempoServico(f *Funcionario) int {
	tempoServico := f.idade - 18
	return tempoServico
}

func main() {
	funcionario := Funcionario{
		nome:    "João",
		salario: 5000.0,
		idade:   30,
	}

	fmt.Println("Salário inicial:", funcionario.salario)
	aumentarSalario(&funcionario, 10)
	fmt.Println("Salário após aumento:", funcionario.salario)
	diminuirSalario(&funcionario, 5)
	fmt.Println("Salário após diminuição:", funcionario.salario)

	tempoServico := calcularTempoServico(&funcionario)
	fmt.Println("Tempo de serviço:", tempoServico, "anos")
}
