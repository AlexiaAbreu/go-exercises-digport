package main

import "fmt"

func comparaEImprimeNumeros(num int) string {
	if num > 0 {
		return "greater than zero"
	} else if num < 0 {
		return "less than zero"
	} else {
		return "zero"
	}
}

func main() {
	var num int
	fmt.Println("Insira um número: ")
	fmt.Scan(&num)
	resultado := comparaEImprimeNumeros(num)
	fmt.Println(resultado)
}
