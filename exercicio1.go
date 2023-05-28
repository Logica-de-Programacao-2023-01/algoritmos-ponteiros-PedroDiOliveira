// Escreva uma função que receba um ponteiro para um inteiro e um valor inteiro n.
// A função deve atualizar o valor do inteiro para a soma dos n primeiros números naturais.
package main

import "fmt"

func somaDosInt(inteiro int, ponteiro *int) int {
	var soma int
	for i := 0; i < inteiro+1; i++ {
		soma += i
	}
	*ponteiro = soma
	return *ponteiro
}

func main() {
	var inteiro int = 5
	var valor int

	ponteiro := &valor

	soma := somaDosInt(inteiro, ponteiro)

	fmt.Printf("A soma dos inteiros de 0 a %d é: %d\n", inteiro, soma)
}
