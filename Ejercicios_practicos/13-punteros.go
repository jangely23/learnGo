/* crea un struc tipo persona,
crea un funcion cambiame con una persona como parametro.
--para des-referenciar un struc, usa (*valor).campo

crea un valor de tipo persona
--imprime el valor
en fun main llamar a cambiame
-- imprime valor
*/

package main

import "fmt"

type persona struct {
	nombre   string
	apellido string
	edad     int
}

func main() {
	p := persona{
		nombre:   "jessica",
		apellido: "leonel",
		edad:     30,
	}

	fmt.Println(p)
	cambiame(&p) //envio la direccin de memoria
	fmt.Println(p)
}

func cambiame(s *persona) {
	s.edad = 31 //tambien se puede hace con (*s).edad
}
