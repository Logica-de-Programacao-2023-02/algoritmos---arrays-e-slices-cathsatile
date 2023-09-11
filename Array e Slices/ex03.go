package main

import "fmt"

func main() {
	num := [4]float64{1.5, 2.5, 3.5, 4.5}
	var mult float64 = 1
	for _, prod := range num {
		mult *= prod
	}
	fmt.Print(mult)
}
