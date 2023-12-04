package main

import "fmt"

func main() {
	fmt.Print("Digite a idade: ")
	var idade int
	fmt.Scanln(&idade)

	var classificacao string

	switch {
	case idade <= 9:
		classificacao = "Mirim"
	case idade >= 10 && idade <= 13:
		classificacao = "Infantil"
	case idade >= 14 && idade <= 17:
		classificacao = "Juvenil"
	default:
		classificacao = "Adulto"
	}

	fmt.Printf("A classificação é: %s\n", classificacao)
}
