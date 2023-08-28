package main

import "fmt"

func main() {

	i := 15
	if i%3 == 0 && i%5 == 0 {
		fmt.Println("i é multiplo")
	} else if i%3 != 0 && i%5 != 0 {
		fmt.Println("i não é multiplo")
	}
}
