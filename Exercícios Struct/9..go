package main

import "fmt"

// Struct Aluno
type Aluno struct {
	nome  string
	idade int
	notas []float64
}

// Função para adicionar uma nota ao aluno
func adicionarNota(a *Aluno, nota float64) {
	a.notas = append(a.notas, nota)
}

// Função para remover uma nota do aluno
func removerNota(a *Aluno, indice int) {
	if indice >= 0 && indice < len(a.notas) {
		a.notas = append(a.notas[:indice], a.notas[indice+1:]...)
	}
}

// Função para calcular a média das notas do aluno
func calcularMediaNotas(a Aluno) float64 {
	soma := 0.0

	for _, nota := range a.notas {
		soma += nota
	}

	media := soma / float64(len(a.notas))
	return media
}

// Função para imprimir o nome, idade e média do aluno
func imprimirInformacoesAluno(a Aluno) {
	fmt.Println("Nome:", a.nome)
	fmt.Println("Idade:", a.idade)
	fmt.Println("Média das notas:", calcularMediaNotas(a))
}

func main() {
	aluno := Aluno{
		nome:  "Maria",
		idade: 20,
		notas: []float64{8.5, 7.2, 9.0},
	}

	imprimirInformacoesAluno(aluno)
	fmt.Println()

	adicionarNota(&aluno, 6.8)
	imprimirInformacoesAluno(aluno)
	fmt.Println()

	removerNota(&aluno, 1)
	imprimirInformacoesAluno(aluno)
}
