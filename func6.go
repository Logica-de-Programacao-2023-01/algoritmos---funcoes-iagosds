package main

import (
	"fmt"
)

// Escreva uma função que receba um slice de strings como parâmetro e retorne uma string com todas as strings
// concatenadas e separadas por vírgulas. Caso o slice esteja vazio, retorne um erro
func virg(sl []string) string {
	frase := ""
	for _, i := range sl {
		frase += i + ","
	}
	return frase
}
func main() {
	var str = []string{"ei", "jutbru", "vai", "tomar", "no", "cu"}
	res := virg(str)
	fmt.Print(res)
}
