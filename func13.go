package main

import "fmt"

//Crie uma função que receba um número inteiro como parâmetro e retorne a soma de todos os seus dígitos.
//Caso o número seja negativo, retorne um erro

func soma(x int) (int, error) {
	sm := 0
	for x > 0 {
		dig := x % 10
		sm += dig
		x /= 10
	}
	if x < 0 {
		return 0, fmt.Errorf("negative value")
	}
	return sm, nil
}

func main() {
	var n int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&n)
	res, err := soma(n)
	if err != nil {
		fmt.Printf("Oorreu um erro: %s\n", err)
	} else {
		fmt.Print("A soma dos caractéres do número é: ", res)
	}
}
