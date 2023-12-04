package main

import "fmt"

func main() {
	var maior, numero int

	fmt.Println("Digite vários números inteiros. Insira 0 para encerrar.")

	for {
		fmt.Print("Digite um número: ")
		fmt.Scanln(&numero)

		if numero == 0 {
			break
		}

		if numero > maior {
			maior = numero
		}
	}

	if maior != 0 {
		fmt.Printf("O maior número é: %d\n", maior)
	} else {
		fmt.Println("Nenhum número foi inserido.")
	}
}
