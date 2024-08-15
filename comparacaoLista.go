package main

import "fmt"

func comparacao(listaDeNum []int) {
	var soma int
	for _, num := range listaDeNum {
		if num > 0 {
			soma = num + 10
			fmt.Printf("%d, ", soma)
		} else if num < 0 {
			soma = num + 23
			fmt.Printf("%d, ", soma)
		} else {
			soma = num + 2
			fmt.Printf("%d, ", soma)
		}
	}
}

func main() {
	listaDeNum := []int{5, 0, -3, 10, -5}
	comparacao(listaDeNum)
}
