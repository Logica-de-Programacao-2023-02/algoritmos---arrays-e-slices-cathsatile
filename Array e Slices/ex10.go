package main

import "fmt"

func main() {
	slice := []int{5, 8, 2, 10, 1, 7, 4, 3, 9, 6}

	var minimo, maximo int
	for _, valor := range slice {
		if valor < minimo {
			minimo = valor
		}
		if valor > maximo {
			maximo = valor
		}
	}

	fmt.Printf("Valor mínimo: %d\n", minimo)
	fmt.Printf("Valor máximo: %d\n", maximo)
}
