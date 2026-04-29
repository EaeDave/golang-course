package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main.")
	auxiliar.Escrever()
	emailValido := checkmail.ValidateFormat("devbook@gmail.com")
	emailInvalido := checkmail.ValidateFormat("123")

	fmt.Println((emailValido))
	fmt.Println(emailInvalido)
}
