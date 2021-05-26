/* tendra un struc de persona con nombre apellido y edad
tendra un metodo para presentarse con el nombre y apellido
crea un valor tipo persona y llama el metodo */

package main

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	persona1 := persona{
		nombre:   "Jessica",
		apellido: "Leonel",
		edad:     31,
	}
	persona1.presentarse()
}

func (p persona) presentarse() {
	println("Mi nombre es:", p.nombre, p.apellido, "y tengo", p.edad, "años")
}
