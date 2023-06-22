package main

import "fmt"

// Crie uma função que receba um número inteiro como parâmetro e retorne um novo slice contendo todos os números primos
// menores ou iguais a ele. Caso o número seja negativo, retorne um erro
func primo(x int) ([]int, error) {
	primos := []int{}
	if x < 0 {
		return primos, fmt.Errorf("Número negativo")
	}
	for i := x; i >= 0; i-- {
		if i == 2 || i == 3 || i == 5 {
			primos = append(primos, i)
		} else if i%2 != 0 && i%3 != 0 && i%5 != 0 && i != 1 {
			primos = append(primos, i)
		}
	}
	return primos, nil
}

func main() {
	num := 120
	res, err := primo(num)
	fmt.Println(res)
	fmt.Println(err)
}
