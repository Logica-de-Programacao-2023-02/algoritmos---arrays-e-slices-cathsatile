package main

import "fmt"

func main() {
	array := [10]int{0, 4, 5, 8, 15, 23, 32, 66, 78, 99}
	var num int
	fmt.Print("Digite o número que queira verificar: ")
	fmt.Scan(&num)
	var verd bool
	for _, elemento := range array {
		if elemento == num {
			verd = true
		}
	}
	if verd {
		fmt.Println("O número está presente no array.")
	} else {
		fmt.Println("O número não está presente no array.")
	}

}
