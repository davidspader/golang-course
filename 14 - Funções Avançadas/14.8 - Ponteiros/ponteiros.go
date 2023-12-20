package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	//USANDO PONTEIRO
	outroNumero := 23
	fmt.Println(outroNumero)
	inverterSinalComPonteiro(&outroNumero)
	fmt.Println(outroNumero)
}
