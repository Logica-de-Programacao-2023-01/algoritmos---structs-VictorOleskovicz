package main

import "fmt"

// Struct Triângulo
type Triangulo struct {
	base   float64
	altura float64
}

// Função para calcular a área do triângulo
func calcularAreaTriangulo(t Triangulo) float64 {
	area := t.base * t.altura / 2
	return area
}

func main() {
	triangulo := Triangulo{base: 5.0, altura: 3.0}
	area := calcularAreaTriangulo(triangulo)
	fmt.Printf("Área do triângulo: %.2f\n", area)
}
