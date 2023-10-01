package main

import "fmt"

func main() {
	array := [10]float64{1, 3.5, 5.6, 7, 5.3, 8.9, 4, 2, 1.5, 6}
	for _, i := range array {
		if i >= 5 {
			slice := []float64{i}
			fmt.Print(slice)
		}
	}
}
