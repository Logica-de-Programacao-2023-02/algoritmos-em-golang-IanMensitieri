package main

import (
	"fmt"
)

func main() {
	fmt.Print("Digite a altura (em metros): ")
	var altura float64
	fmt.Scanln(&altura)

	fmt.Print("Digite o sexo (M para masculino, F para feminino): ")
	var sexo string
	fmt.Scanln(&sexo)

	var pesoIdeal float64
	const fatorIMC = 22.0 // Peso ideal para um IMC de 22.0

	if sexo == "M" || sexo == "m" {
		pesoIdeal = fatorIMC * (altura * altura)
	} else if sexo == "F" || sexo == "f" {
		pesoIdeal = fatorIMC * (altura * altura)
	} else {
		fmt.Println("Sexo inválido. Use M ou F.")
		return
	}

	fmt.Printf("O peso ideal para a altura %.2f metros e sexo %s é: %.2f kg\n", altura, sexo, pesoIdeal)
}
