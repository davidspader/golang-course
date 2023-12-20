package main

import "fmt"

func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	fmt.Println(soma(123, 12, 321, 3, 423, 213, 4, 123, 123, 24, 12, 312, 2))

	escrever("O numero é: ", 3, 213, 213, 213, 213, 21, 321, 321, 3, 213, 21)
}
