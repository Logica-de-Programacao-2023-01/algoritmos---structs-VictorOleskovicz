package main

import "fmt"

// Struct Musica
type Musica struct {
	titulo   string
	artista  string
	duracao  int
}

// Struct Playlist
type Playlist struct {
	nome    string
	musicas []Musica
}

// Função para imprimir informações da playlist
func imprimirPlaylist(p Playlist) {
	fmt.Println("Playlist:", p.nome)

	totalDuracao := 0

	for _, musica := range p.musicas {
		fmt.Printf("Título: %s, Artista: %s, Duração: %d segundos\n", musica.titulo, musica.artista, musica.duracao)
		totalDuracao += musica.duracao
	}

	fmt.Printf("Tempo total da playlist: %d segundos\n", totalDuracao)
}

func main() {
	musicas := []Musica{
