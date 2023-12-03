package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Digite um número inteiro positivo n: ")
	fmt.Scan(&n)

	var primos []int
	num := 2

	for len(primos) < n {
		primo := true
		limite := num / 2

		for i := 2; i <= limite; i++ {
			if num%i == 0 {
				primo = false
				break
			}
		}
		if primo {
			primos = append(primos, num)
		}
		num++
	}

	fmt.Printf("Os %d primeiros números primos são: %v\n", n, primos)
}
