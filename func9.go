package main

import "fmt"

//Crie uma função que receba uma string como parâmetro e retorne um novo slice com todas as palavras contidas na string.
//Considere que as palavras são separadas por espaços em branco. Caso a string seja vazia, retorne um erro

func frase(palavras []string) (string, error) {
	nova := ""
	for _, i := range palavras {
		nova += " " + i
	}
	if nova == nil {
		return nova, fmt.Errorf("Slice vazia")
	}
	return nova, nil
}
func main() {
	var slice = []string{"gol", "do", "fluminense"}
	res := frase(slice)
	fmt.Println(res)
}
