package main

import "fmt"

//Escreva uma função que receba dois slices de inteiros como parâmetros e retorne um novo slice contendo apenas os
//números presentes nos dois slices. Caso um dos slices esteja vazio, retorne um erro

func uniao(x []int, y []int) ([]int, error) {
	var union []int
	if x == nil || y == nil {
		return union, fmt.Errorf("Algum dos slices esta vazio")
	}
	for _, i := range x {
		exist := false
		for _, n := range union {
			if i == n {
				exist = true
			}
		}
		if exist {
			continue
		}
		for _, c := range y {
			exist := false
			for _, n := range union {
				if c == n {
					exist = true
				}
			}
			if exist {
				continue
			}
			if i == c {
				union = append(union, i)
			}
		}
	}
	if union == nil {
		return union, fmt.Errorf("Slice de retorno esta vazio")
	}
	return union, nil
}
func main() {
	var p = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var t = []int{1, 3, 3, 5, 7, 7, 9}
	res, err := uniao(p, t)
	fmt.Println(res)
	fmt.Println(err)
}
