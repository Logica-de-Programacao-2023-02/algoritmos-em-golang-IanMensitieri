package main

import (
	"fmt"
)

func main() {
	fmt.Print("Digite o peso em quilos: ")
	var pesoEmKilos float64
	fmt.Scanln(&pesoEmKilos)

	taxaConversao := 2.20462
	pesoEmLibras := pesoEmKilos * taxaConversao
	fmt.Printf("O peso em libras é: %.2f\n", pesoEmLibras)
}
