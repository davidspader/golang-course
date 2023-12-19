package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	var u usuario
	u.nome = "David"
	u.idade = 26
	fmt.Println(u)

	enderecoExeplo := endereco{"blabalbla", 2}

	usuario2 := usuario{"David2", 20, enderecoExeplo}
	fmt.Println(usuario2)

	usuario3 := usuario{idade: 20}
	fmt.Println(usuario3)
}
