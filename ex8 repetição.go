package main

import "fmt"

func main() {
	fmt.Print("Digite um número inteiro positivo: ")
	var numero int
	fmt.Scanln(&numero)

	if numero <= 0 {
		fmt.Println("Por favor, digite um número inteiro positivo.")
		return
	}

	fmt.Printf("Os divisores de %d são: ", numero)
	for i := 1; i <= numero; i++ {
		if numero%i == 0 {
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()
}
