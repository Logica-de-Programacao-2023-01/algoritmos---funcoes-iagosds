package main

import "fmt"

// Crie uma função que receba um slice de inteiros e retorne o segundo maior valor
func segundo(slice []int) int {
	maior := slice[0]
	sec := 0
	for i := range slice {
		if slice[i] > maior {
			sec = maior
			maior = slice[i]
		}
	}
	return sec
}
func main() {
	var sl = []int{1, 2, 3, 4, 2, 1, 5}
	res := segundo(sl)
	fmt.Print(res)
}
