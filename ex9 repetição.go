package main

import "fmt"

func main() {
	var total, quantidade int
	var numero float64

	fmt.Println("Digite vários números inteiros. Insira 0 para encerrar.")

	for {
		fmt.Print("Digite um número: ")
		fmt.Scanln(&numero)

		if numero == 0 {
			break
		}

		total += int(numero)
		quantidade++
	}

	var media float64
	if quantidade > 0 {
		media = float64(total) / float64(quantidade)
		fmt.Printf("A média aritmética é: %.2f\n", media)
	} else {
		fmt.Println("Nenhum número foi inserido.")
	}
}
