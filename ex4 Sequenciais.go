package main

import "fmt"

func main() {

	peso := 60.
	altura := 1.70
	imc := peso / (altura * altura)
	if imc < 18.5 {
		fmt.Println("abaixo do peso normal")
	} else if imc >= 18.5 && imc < 24.9 {
		fmt.Println("peso normal")
	} else if imc >= 25 && imc < 29.9 {
		fmt.Println("Obesidade classe 1")
	}
}
