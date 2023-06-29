package main

import (
	"fmt"
	"math"
)

// Struct Circulo
type Circulo struct {
	raio float64
}

// Função para calcular a área do círculo
func calcularAreaCirculo(c Circulo) float64 {
	area := math.Pi * c.raio * c.raio
	return area
}

func main() {
	circulo := Circulo{raio: 3.0}
	area := calcularAreaCirculo(circulo)
	fmt.Printf("Área do círculo: %.2f\n", area)
}
