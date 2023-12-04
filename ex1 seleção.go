package main

import (
	"fmt"
)

func main() {
	// Solicitar os dois números inteiros ao usuário
	fmt.Print("Digite o primeiro número inteiro: ")
	var num1 int
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número inteiro: ")
	var num2 int
	fmt.Scanln(&num2)

	// Verificar qual número é maior
	var maior int
	if num1 > num2 {
		maior = num1
	} else {
		maior = num2
	}

	// Exibir o resultado
	fmt.Printf("O maior número é: %d\n", maior)
}
