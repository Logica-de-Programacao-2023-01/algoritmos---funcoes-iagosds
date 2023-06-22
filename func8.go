package main

import "fmt"

// Escreva uma função que receba um slice de inteiros como parâmetro e retorne um novo slice com apenas os números pares
// contidos no slice. Caso o slice esteja vazio, retorne um erro
func pares(sl []int) ([]int, error) {
	var novo []int
	for _, i := range sl {
		if i%2 == 0 {
			novo = append(novo, i)
		}
	}
	if novo == nil {
		return novo, fmt.Errorf("Slice vazio")
	}
	return novo, nil
}
func main() {
	var slice = []int{1, 3, 5, 7}
	res, err := pares(slice)
	fmt.Println(res)
	fmt.Println(err)
}
