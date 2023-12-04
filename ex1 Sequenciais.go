package main

import "fmt"

func main() {

	x := 9
	y := 11
	if x > y {
		fmt.Println("x é maior que y")
	} else if x > y {
		fmt.Println("y é maior que x")
	} else {
		fmt.Println("x e y são iguais")
	}
}
