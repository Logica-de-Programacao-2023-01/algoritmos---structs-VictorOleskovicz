package main

import "fmt"

// Struct Filme
type Filme struct {
	titulo     string
	diretor    string
	ano        int
	avaliacoes []int
}

// Função para adicionar uma avaliação ao filme
func adicionarAvaliacao(f *Filme, avaliacao int) {
	f.avaliacoes = append(f.avaliacoes, avaliacao)
}

// Função para remover uma avaliação do filme
func removerAvaliacao(f *Filme, indice int) {
	if indice >= 0 && indice < len(f.avaliacoes) {
		f.avaliacoes = append(f.avaliacoes[:indice], f.avaliacoes[indice+1:]...)
	}
}

// Função para calcular a média das avaliações do filme
func calcularMediaAvaliacoes(f Filme) float64 {
	soma := 0

	for _, avaliacao := range f.avaliacoes {
		soma += avaliacao
	}

	media := float64(soma) / float64(len(f.avaliacoes))
	return media
}

// Função para imprimir as informações do filme e sua média de avaliações
func imprimirInformacoesFilme(f Filme) {
	fmt.Println("Título:", f.titulo)
	fmt.Println("Diretor:", f.diretor)
	fmt.Println("Ano:", f.ano)
	fmt.Println("Média das avaliações:", calcularMediaAvaliacoes(f))
}

func main() {
	filme := Filme{
		titulo:     "O Poderoso Chefão",
		diretor:    "Francis Ford Coppola",
		ano:        1972,

