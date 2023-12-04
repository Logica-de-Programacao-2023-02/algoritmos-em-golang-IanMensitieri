package main

import "fmt"

func main() {
	fmt.Print("Digite o primeiro número inteiro: ")
	var num1 int
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número inteiro: ")
	var num2 int
	fmt.Scanln(&num2)

	if num1 >= 0 && num2 >= 0 {
		resultado := num1 * num2
		fmt.Printf("O resultado da multiplicação é: %d\n", resultado)
	} else {
		resultado := num1 + num2
		fmt.Printf("O resultado da soma é: %d\n", resultado)
	}
}
