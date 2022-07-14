/*
As funções Go podem ser escritas para funcionar em vários
tipos usando parâmetros de tipo. Os parâmetros de tipo de uma função aparecem entre colchetes,
antes dos argumentos da função.

	func Index[T comparable](s []T, x T) int

Esta declaração significa que s é um slice de qualquer tipo T que cumpre o restrição interna comparable.
x também é um valor do mesmo tipo.

comparable é uma restrição útil que torna possível usar os operadores == e != em valores do tipo.
Neste exemplo, usamos para comparar um valor para todos os elementos da slice até que uma correspondência
seja encontrada. Esta função Index funciona para qualquer tipo que suporte comparação.

*/

package main

import "fmt"

// Index returns the index of x in s, or -1 if not found.
func Index[T comparable](s []T, x T) int {
	for i, v := range s {
		// v and x are type T, which has the comparable
		// constraint, so we can use == here.
		if v == x {
			return i
		}
	}
	return -1
}

func main() {
	// Index works on a slice of ints
	si := []int{10, 20, 15, -10}
	fmt.Println(Index(si, 15))

	// Index also works on a slice of strings
	ss := []string{"foo", "bar", "baz"}
	fmt.Println(Index(ss, "hello"))
}
