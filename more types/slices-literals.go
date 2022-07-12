// Uma slice literal é como um array literal, sem o comprimento.

package main

import "fmt"

func main() {
	q := []int{2, 3, 5, 7, 11, 13}
	fmt.Println(q)

	r := []bool{true, false, true, true, false, true}
	fmt.Println(r)

	s := []struct {
		i int
		b bool
	}{
		{2, true},
		{3, true},
		{5, true},
		{7, true},
		{11, true},
		{13, true},
		{6, false},
	}
	fmt.Println(s)
}
