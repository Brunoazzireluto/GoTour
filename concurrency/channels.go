/*
Canais são um conduto tipado através do qual você pode
enviar e receber valores com o operador de canal, <-.

	ch <- v    // v envia para o canal ch.
	v := <-ch  // Recebe do ch, e
			// atribui o valor de v.
	(Os dados fluem na direção da seta.)

Como maps e slices, os canais devem ser criados antes de se usar:

	ch := make(chan int)
Por padrão, enviam e recebem bloco até o outro lado estar pronto.
Isso permite que goroutines sincronizem sem bloqueios explícitos ou variáveis ​​de condição.

O código do exemplo soma os números em uma slice,
distribuindo o trabalho entre duas goroutines.
Uma vez que ambas as goroutines completaram seu processamento,
calcula o resultado final.
*/

package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)

	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c
	fmt.Println(x, y, x+y)
}
