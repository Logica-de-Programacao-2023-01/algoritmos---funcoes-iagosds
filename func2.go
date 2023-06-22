package main

import (
	"fmt"
	"strings"
)

// Crie uma função que receba uma string e retorne a quantidade de vogais presentes
func vogais(frase string) int {
	v := "aeiouAEIOU"
	cont := 0
	for _, i := range frase {
		if strings.ContainsRune(v, i) {
			cont++
		}
	}
	return cont
}
func main() {
	str := "jutbruta"
	res := vogais(str)
	fmt.Print(res)
}
