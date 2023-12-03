package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8}
	var num, num2 int
	fmt.Print("Digite o índice do primeiro elemento: ")
	fmt.Scan(&num)
	fmt.Print("Digite o índice do segundo elemento: ")
	fmt.Scan(&num2)
	slice[num], slice[num2] = slice[num2], slice[num]
	fmt.Println("A nova slice fica: ", slice)

}
