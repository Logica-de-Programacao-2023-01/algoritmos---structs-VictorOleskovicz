package main

import "fmt"

// Struct Musica
type Musica struct {
	titulo  string
	artista string
	duracao int
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

// Função para buscar playlists por título de música
func buscarPlaylistsPorTitulo(titulo string, playlists []Playlist) []Playlist {
	var playlistsEncontradas []Playlist

	for _, playlist := range playlists {
		for _, musica := range playlist.musicas {
			if musica.titulo == titulo {
				playlistsEncontradas = append(playlistsEncontradas, playlist)
				break
			}
		}
	}

	return playlistsEncontradas
}

func main() {
	playlist1 := Playlist{
		nome: "Playlist 1",
		musicas: []Musica{
			{titulo: "Música 1", artista: "Artista 1", duracao: 180},
			{titulo: "Música 2", artista: "Artista 2", duracao: 240},
		},
	}

	playlist2 := Playlist{
		nome: "Playlist 2",
		musicas: []Musica{
			{titulo: "Música 3", artista: "Artista 3", duracao: 300},
			{titulo: "Música 4", artista: "Artista 4", duracao: 200},
		},
	}

	playlist3 := Playlist{
		nome: "Playlist 3",
		musicas: []Musica{
			{titulo: "Música 1", artista: "Artista 1", duracao: 180},
			{titulo: "Música 5", artista: "Artista 5", duracao: 210},
		},
	}

	playlists := []Playlist{playlist1, playlist2, playlist3}

	tituloBuscado := "Música 1"
	playlistsEncontradas := buscarPlaylistsPorTitulo(tituloBuscado, playlists)

	fmt.Printf("Playlists que possuem a música '%s':\n", tituloBuscado)
	for _, playlist := range playlistsEncontradas {
		fmt.Println(playlist.nome)
	}
}
