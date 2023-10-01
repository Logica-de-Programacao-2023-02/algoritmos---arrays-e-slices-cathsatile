package main

import (
	"fmt"
	"log"
)

func main() {
	array := [6]float64{1.2, 1.3, 1.4, 1.5, 1.6, 1.7}

	var x float64
	fmt.Print("Digite um número: ")
	fmt.Scan(&x)

	for i := 0; i < len(array); i++ {
		array[i] += x
	}
	log.Print(array)
}
