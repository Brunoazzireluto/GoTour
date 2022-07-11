// Este grupo de códigos em parênteses da importação, é a declaração de importação "consignada". Você também pode escrever várias declarações de importação assim:

// import "fmt"
// import "math"
// Mas é um bom estilo usar a declaração de importação consignada.

package gotour

import (
	"fmt"
	"math"
)

func main() {
	fmt.Printf("Agora você tem %g problemas. \n", math.Sqrt(7))
}
