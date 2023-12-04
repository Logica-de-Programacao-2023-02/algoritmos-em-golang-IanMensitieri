package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Print("Digite o primeiro número real: ")
	var num1 float64
	fmt.Scanln(&num1)

	fmt.Print("Digite o segundo número real: ")
	var num2 float64
	fmt.Scanln(&num2)

	fmt.Print("Digite o terceiro número real: ")
	var num3 float64
	fmt.Scanln(&num3)

	numeros := []float64{num1, num2, num3}
	sort.Float64s(numeros)

	fmt.Println("Os números em ordem crescente são:", numeros)
}
