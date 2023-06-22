package main

import "fmt"

//Crie uma função que receba um slice de inteiros e retorne a média aritmética dos valores
func media(sl []int) float64 {
	var media float64 = 0
	for _, i := range sl {
		media += float64(i)
	}
	media /= float64(len(sl))
	return media
}
func main() {
	var slice = []int{2, 3, 4, 1}
	res := media(slice)
	fmt.Print(res)

}
