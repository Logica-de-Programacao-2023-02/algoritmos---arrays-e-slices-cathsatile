package main

import "fmt"

func main() {
	var tamanho int
	fmt.Print("Digite o tamanho da slice: ")
	fmt.Scan(&tamanho)

	slice := make([]int, tamanho)
	for i := 0; i < tamanho; i++ {
		fmt.Print("Digite os números: ")
		fmt.Scan(&slice[i])
	}
	var soma int
	for _, num := range slice {
		soma += num
	}
	fmt.Println("slice [", tamanho, "]={", slice, "}")
	fmt.Println("A soma dos números da slice é: ", soma)
}
