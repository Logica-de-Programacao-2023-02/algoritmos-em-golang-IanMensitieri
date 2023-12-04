package main

import (
	"fmt"
)

func main() {
	fmt.Print("Digite o primeiro número inteiro: ")
	var num1 int
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número inteiro: ")
	var num2 int
	fmt.Scanln(&num2)

	fmt.Print("Digite o terceiro número inteiro: ")
	var num3 int
	fmt.Scanln(&num3)
	var menor int
	if num1 < num2 && num1 < num3 {
		menor = num1
	} else if num2 < num3 {
		menor = num2
	} else {
		menor = num3
	}

	fmt.Printf("O menor número é: %d\n", menor)
}
