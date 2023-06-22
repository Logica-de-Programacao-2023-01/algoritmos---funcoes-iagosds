package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//Escreva uma função que receba uma string como parâmetro e retorne uma nova string com todas as vogais substituídas
//por "*". Caso a string seja vazia, retorne um erro

func subs(str string) string {
	vowels := "aeiouAEIOU"
	replaced := ""

	for _, c := range str {
		if strings.ContainsRune(vowels, c) {
			replaced += "*"
		} else {
			replaced += string(c)
		}
	}

	return replaced
}

func main() {

	fmt.Print("Digite uma string: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	f := scanner.Text()
	res := subs(f)

	fmt.Print("String com as substituições: ", res)
}
