package main

import "fmt"

//Crie uma função que receba um slice de strings e retorne a concatenação de todas as strings

func concat(vs []string) string {
	var c string
	for i := 0; i < len(vs); i++ {
		fmt.Print("Digite a ", i+1, "ª palavra do vetor: ")
		fmt.Scanln(&vs[i])
		c += vs[i] + " "
	}
	return c
}

func main() {
	var n int
	fmt.Print("Quantas palavras deseja armazenar no vetor? ")
	fmt.Scanln(&n)
	var sl []string = make([]string, n)
	fmt.Println("Resultado", concat(sl))
}
