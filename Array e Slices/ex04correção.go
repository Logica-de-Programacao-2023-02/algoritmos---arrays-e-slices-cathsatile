package main

import "fmt"

func main() {
	var tamanho, valor int
	fmt.Print("Digite o tamanho da slice: ")
	fmt.Scan(&tamanho)
	slice := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Print("Digite o valor: ")
		fmt.Scan(&valor)
		slice[i] = valor
	}
	var soma int
	for i := 0; i < len(slice); i++ {
		x := slice[i]
		soma += x
	}
	fmt.Println("Meu slice é: ", slice)
	fmt.Println("A soma é: ", soma)
}
