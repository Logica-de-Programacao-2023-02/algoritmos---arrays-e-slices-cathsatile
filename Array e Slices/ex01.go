package main

import "fmt"

func main() {
	var numeros = [3]int{1, 2, 3}
	var soma int
	for _, num := range numeros {
		soma += num
	}
	fmt.Println(soma)

}
