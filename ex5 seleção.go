package main

import "fmt"

func main() {
	fmt.Print("Digite um número inteiro: ")
	var numero int
	fmt.Scanln(&numero)

	if numero%3 == 0 && numero%5 == 0 {
		fmt.Printf("%d é múltiplo de 3 e de 5.\n", numero)
	} else {
		fmt.Printf("%d não é múltiplo de 3 e de 5 ao mesmo tempo.\n", numero)
	}
}
