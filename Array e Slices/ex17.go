package main

import "fmt"

func main() {
	array := [10]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var soma int
	for _, i := range array {
		if i%2 == 0 {
			soma += i
		}
	}
	fmt.Print(soma)

}
