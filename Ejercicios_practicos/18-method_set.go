/*crear un struct tipo persona que tenga el metodo hablar.
crear la interface humano para usar implicitamente el metodo hablar
crear la funcion diAlgo que tome un humano como parametro y llame el metodo hablar
*/

package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

type humano interface {
	hablar()
}

func main() {
	p1 := persona{
		nombre:   "Jessica",
		apellido: "Leonel",
	}

	p2 := persona{
		nombre:   "Lauren",
		apellido: "Laguna",
	}

	diAlgo(&p1)
	diAlgo(&p2)
	p1.hablar() //al llamar el metodo directamente se utiliza la direccion de memoria
}

func (r *persona) hablar() {
	fmt.Println("Hola, soy: ", r.nombre, r.apellido, " ¿como estas?")
}

func diAlgo(h humano) {
	h.hablar()
}
