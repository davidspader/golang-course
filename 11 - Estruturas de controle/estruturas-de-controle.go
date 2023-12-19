package main

import "fmt"

func main() {
	numero := -2

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("O número é maior que zero")
	} else if outroNumero < -10 {
		fmt.Println("O número é menor que -10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
