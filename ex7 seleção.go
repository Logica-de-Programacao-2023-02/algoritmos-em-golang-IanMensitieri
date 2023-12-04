package main

import "fmt"

func main() {
	fmt.Print("Digite o salário do funcionário: ")
	var salario float64
	fmt.Scanln(&salario)

	const taxaAumentoMenor1000 = 0.10  // 10% de aumento para salários abaixo de R$ 1000,00
	const taxaAumentoMaiorOuIgual1000 = 0.05  // 5% de aumento para salários R$ 1000,00 ou acima

	var novoSalario float64
	if salario < 1000.00 {
		novoSalario = salario * (1 + taxaAumentoMenor1000)
	} else {
		novoSalario = salario * (1 + taxaAumentoMaiorOuIgual1000)
	}

	fmt.Printf("O novo salário é: R$ %.2f\n", novoSalario)
}
