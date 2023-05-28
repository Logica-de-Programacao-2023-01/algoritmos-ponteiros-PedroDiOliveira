// Escreva uma função que receba um ponteiro para um inteiro e verifique se esse inteiro é par ou ímpar.
// A função deve atualizar o valor do inteiro para 0 se ele for par ou para 1 se for ímpar.
package main

import "fmt"

func ptr(p *int) {
	if *p%2 == 0 {
		*p = 0
	} else if *p%2 != 0 {
		*p = 1
	}
}

func main() {
	var num int
	num = 95
	ptr(&num)
	fmt.Println(num)
}
