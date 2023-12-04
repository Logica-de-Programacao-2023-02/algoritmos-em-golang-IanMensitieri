package main

import "fmt"

func main() {
	fmt.Print("Digite um número inteiro: ")
	var numero int
	fmt.Scanln(&numero)

	if numero%2 == 0 {
		fmt.Println("O número é par.")
	} else {
		fmt.Println("O número é ímpar.")
	}
}
