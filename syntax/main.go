package main

import "fmt"

func main() {
	slices := generateSlices()

	for _, s := range slices {
		if s%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}

func generateSlices() []int {
	slices := []int{}
	for i := 0; i <= 10; i++ {
		slices = append(slices, i)
	}

	return slices
}
