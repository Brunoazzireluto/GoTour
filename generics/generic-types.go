/*
Além das funções genéricas,
Go também oferece suporte a tipos genéricos. Um tipo pode ser parametrizado com um parâmetro de tipo,
que pode ser útil para implementar estruturas de dados genéricas.

Este exemplo demonstra uma declaração de tipo simples para uma lista vinculada simples
segurando qualquer tipo de valor.
*/

package main

// List represents a singly-linked list that holds
// values of any type.
type List[T any] struct {
	next *List[T]
	val  T
}

func main() {
}
