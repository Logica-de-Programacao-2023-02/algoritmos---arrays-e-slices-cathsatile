package main

import "fmt"

func main() {
	array := []int{1, 2, 3, 4, 5, 6, 7}
	var num, num2 int
	fmt.Print("Adicione o número do começo do array: ")
	fmt.Scan(&num)
	fmt.Print("Adicione o número do final da array: ")
	fmt.Scan(&num2)
	array = append([]int{num}, array...)
	array = append(array, num2)
	fmt.Println("O novo array fica: ", array)
}
