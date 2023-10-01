package main

import "fmt"

func main() {
	array := []int{3, 5, 18, 21, 31}
	for _, i := range array {
		if i%3 == 0 {
			slice := []int{i}
			fmt.Print(slice)
		}

	}
}
