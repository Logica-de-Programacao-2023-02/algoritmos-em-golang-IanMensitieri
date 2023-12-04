package main

import "fmt"

func main() {

	x := 4
	y := 5
	z := 7
	if x < y && x < z {
		fmt.Println("x é menor q y e z")
	} else if y < x && y < z {
		fmt.Println("y é menor q x e z")
	} else if z < x && z < y {
		fmt.Println("z é menor q x e y")
	}

}
