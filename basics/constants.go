package main

import "fmt"

//Constantes não podem ser declaradas usando a sintaxe :=

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("Hello", World)
	fmt.Println("happy", Pi, "day")

	const Truth = true
	fmt.Println("go rules? ", Truth)
}
