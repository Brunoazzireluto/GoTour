// A struct é uma coleção de campos.

package main

import "fmt"

type Vertex struct {
	x int
	y int
	z string
}

func main() {
	fmt.Println(Vertex{1, 2, "abc"})
}
