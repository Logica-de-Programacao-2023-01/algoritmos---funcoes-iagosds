package main

import (
	"fmt"
)

func ord(s []int) ([]int, error) {
	var p, sg, aux int = 0, 0, 0
	for p = 0; p < len(s); p++ {
		for sg = p + 1; sg < len(s); sg++ {
			if s[p] > s[sg] {
				aux = s[p]
				s[p] = s[sg]
				s[sg] = aux
			}
		}
	}
	if s == nil || len(s) == 0 {
		return s, fmt.Errorf("empty slice!")
	}
	return s, nil
}
func main() {
	var n int
	fmt.Print("Quantos números deseja ordernar? ")
	fmt.Scanln(&n)
	var sl []int = make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Print("Digite o ", i+1, "º número: ")
		fmt.Scan(&sl[i])
	}
	res, err := ord(sl)
	if err != nil {
		fmt.Printf("Ocorreu um erro: %s\n", err)
	} else {
		fmt.Printf("\nA ordem crescente dos elementos é:\n")
		for i := 0; i < len(res); i++ {
			fmt.Print("[", res[i], "] ")
		}
	}
}
