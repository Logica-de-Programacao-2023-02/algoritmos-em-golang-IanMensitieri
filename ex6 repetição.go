package main

import "fmt"

func main() {
	fmt.Print("Digite um número para a tabuada: ")
	var numero int
	fmt.Scanln(&numero)

	fmt.Printf("Tabuada de multiplicação para %d:\n", numero)
	for i := 1; i <= 10; i++ {
		resultado := numero * i
		fmt.Printf("%d x %d = %d\n", numero, i, resultado)
	}
}
