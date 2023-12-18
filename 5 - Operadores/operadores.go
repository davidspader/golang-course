package main

import "fmt"

func main() {
	//ARITIMETICOS
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 40 / 2
	multiplicacao := 1 * 2
	mod := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, mod)

	var numero1 int16 = 10
	var numero2 int16 = 12
	soma2 := numero1 + numero2
	fmt.Println(soma2)

	//ATRIBUIÇÃO
	var variavel1 string = "teste"
	variavel2 := "teste2"
	fmt.Println(variavel1, variavel2)

	//OPERADORES RELACIONAIS
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 != 2)

	//OPERADORES LOGICOS
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	//OPERADORES UNÁRIOS
	numero := 10
	numero++
	numero += 15
	fmt.Println(numero)

	numero--
	numero -= 20
	fmt.Println(numero)

	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	//NÃO EXISTE OPERADOR TERNARIO
}
