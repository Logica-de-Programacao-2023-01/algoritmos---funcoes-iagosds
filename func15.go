package main

import "fmt"

//Crie uma função que receba um número inteiro e um slice de inteiros como parâmetros e retorne verdadeiro se o número
//estiver presente no slice e falso caso contrário. Caso o slice esteja vazio, retorne um erro

func check(x int, lis []int) (bool, error) {
	if lis == nil {
		return false, fmt.Errorf("Slice vazio")
	}
	for i := range lis {
		if lis[i] == x {
			return true, nil
			break
		}
	}

	return false, nil
}
func main() {
	lista := []int{1, 2, 3, 4, 5, 6}
	ele := 7
	res, err := check(ele, lista)
	fmt.Println(res)
	fmt.Println(err)
}
