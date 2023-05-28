// Escreva uma função em Go que receba um ponteiro para um struct Conta com campos saldo e titular,
// e adicione um valor ao saldo da conta.
package main

import "fmt"

type Conta struct {
	Titular string
	Saldo   float64
}

func main() {
	contaBanco := Conta{
		Titular: "Gabriel",
		Saldo:   12.50,
	}
	fmt.Println("Antes do pix:", contaBanco.Saldo)

	banco(&contaBanco, 120)
	fmt.Println("Depois do pix:", contaBanco.Saldo)
}

func banco(c *Conta, pix float64) {
	c.Saldo += pix
}
