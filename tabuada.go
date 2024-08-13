package main

import "fmt"

func multiplica(num int) {
	for i := 0; i <= 10; i++ {
		multiplicacao := i * num
		fmt.Printf("%d x %d = %d\n", i, num, multiplicacao)
	}

}
func main() {
	var input int
	fmt.Println("Digite um número: ")
	fmt.Scan(&input)
	multiplica(input)
}
