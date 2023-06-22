package main

import "fmt"

// Crie uma função que receba um número inteiro como parâmetro e retorne verdadeiro se ele for um número primo e falso
// caso contrário. Caso o número seja negativo, retorne um erro.
func primo(x int) bool {
	if x < 0 {
		return false
	} else if x == 2 || x == 3 {
		return true
	} else if x%2 == 0 || x%3 == 0 || x == 1 || x == 0 {
		return false
	} else {
		return true
	}
	return true
}
func main() {
	var n int
	fmt.Print("Digite um número: ")
	fmt.Scan(&n)
	res := primo(n)
	fmt.Println(res)
}
