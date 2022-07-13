/*
O package image define a interface Image:

	type Image interface {
		ColorModel() color.Model
		Bounds() Rectangle
		At(x, y int) color.Color
	}
Nota: o Rectangle retorna valores do método Bounds que na realidade é um image.Rectangle,
como na declaração está dentro do pacote image.

(Veja a documentação para todos os detalhes)

Os tipos color.Color e color.Model são interfaces também,
mas nós vamos ignorar isso usando implementações pré-definidas color.RGBA e color.RGBAModel.
Essas interfaces e tipo são especificados pelo pacote image/color
*/

package main

import (
	"fmt"
	"image"
)

func main() {
	m := image.NewRGBA(image.Rect(0, 0, 100, 100))
	fmt.Println(m.Bounds())
	fmt.Println(m.At(0, 0).RGBA())
}
