package main

import "fmt"

func main() {
	matriz := [2][3]int{{1, 2, 3}, {4, 5, 6}}
	for i := 0; i <= 2; i++ {
		for j := 0; j <= 3; j++ {
			fmt.Print("Digite a linha: ")
			fmt.Scan(&i)
			fmt.Print("Digite a coluna: ")
			fmt.Scan(&j)
			fmt.Printf("Se a posição for [%d][%d] elemento é: %d\n", i, j, matriz[i][j])
		}
	}
}
