package main

import "fmt"

func add(x int, y int) int {
	return x + y
}

// Quando dois ou mais parâmetros nomeados consecutivos de uma função compartilham o mesmo tipo,
// você pode omitir o tipo de todos menos o último. Neste exemplo, vamos encurtar

func add_shorten(x, y int) int {
	return x + y
}

func main() {
	fmt.Println((add(42, 13)))
	fmt.Println(add(add(55, 25), add(12, 9)))
	fmt.Println(add_shorten(25, 25))
}
