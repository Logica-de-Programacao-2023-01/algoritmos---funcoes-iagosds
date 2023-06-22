package main

import "fmt"

//Escreva uma função que receba um slice de strings como parâmetro e retorne um novo slice contendo apenas as strings
//que possuem mais de 5 caracteres. Caso o slice esteja vazio, retorne um erro

func cinmais(strs []string) ([]string, error) {
	nova := []string{}
	for _, i := range strs {
		if len(i) > 5 {
			nova = append(nova, i+",")
		}
	}
	if nova == nil {
		return nova, fmt.Errorf("Slice vazio")
	}
	return nova, nil
}
func main() {
	frases := []string{"012345", "mais de 5", "-d5", "nao", "claro que sim"}
	res, err := cinmais(frases)
	fmt.Println(res)
	fmt.Println(err)
}
