package main

import "fmt"

func main() {
	num := []int{1, 2, 3, 4, 5}
	var n1 int
	fmt.Print("Digite o número:")
	fmt.Scan(&n1)
	var ver bool
	for i := range num {
		if i == n1 {
			ver = true
		}
	}
	if ver {
		fmt.Println("O número já está presente no slice")

	} else {
		num2 := append(num, n1)
		fmt.Printf("O slice fica: %d\n", num2)
	}
}
