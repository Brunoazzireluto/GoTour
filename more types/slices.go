/* Uma matriz tem um tamanho fixo. Uma slice, por outro lado, é dinamicamente redimensionada,
uma visão flexível dos elementos de uma matriz. Na prática, as slices são muito mais comuns do que as matrizes.
*/

package main

import "fmt"

func main() {
	primes := [6]int{2, 3, 5, 7, 11, 13}

	var s []int = primes[1:4]
	fmt.Println(s)
}
