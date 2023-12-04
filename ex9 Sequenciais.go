package main

import (
	"fmt"
)

func main() {
	fmt.Print("Digite o preço do produto: ")
	var preco float64
	fmt.Scanln(&preco)

	desconto := 0.10
	valorComDesconto := preco * (1 - desconto)

	fmt.Printf("O valor com desconto de 10%% é: %.2f\n", valorComDesconto)
}
