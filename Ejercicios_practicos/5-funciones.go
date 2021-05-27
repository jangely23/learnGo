/* crear un tipo cuadrado y circulo, adjuntarles un metodo area.
el metodo debe retornar el area, debe tener un tipo forma que define la interfaz
crear una funcion info que reciba  un tipo forma y calcule el area y la retorne
finalemente crear dos valores y mostrar el area
*/
package main

import (
	"fmt"
	"math"
)

type circulo struct {
	radio float64
}

type cuadrado struct {
	lado float64
}

type forma interface {
	area() float64
}

func main() {
	cua := cuadrado{
		lado: 15,
	}

	red := circulo{
		radio: 12.345,
	}

	info(cua)
	info(red)

}

func (c circulo) area() float64 {
	return math.Pi * c.radio * c.radio
}

func (c cuadrado) area() float64 {
	return c.lado * c.lado
}

func info(f forma) {

	fmt.Println("El area es:", f.area())
}
