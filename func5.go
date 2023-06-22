package main

import "fmt"

// Crie uma função que receba um slice de inteiros e um valor inteiro, e retorne a posição do primeiro elemento igual ao
// valor no slice. Caso não encontre, retorne -1
func position(slice []int, x int) int {
	pos := -1
	for c, i := range slice {
		if x == i {
			pos = c
			break
		}
	}
	return pos
}
func main() {
	var sl = []int{1, 2, 3, 4, 5}
	res := position(sl, 6)
	fmt.Print(res)
}
