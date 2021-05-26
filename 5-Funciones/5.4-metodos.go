/* Se usa para decirle al programa que ejecute esa funcion al final de la funcion en la que se declaro, ¿como de termina el sistema que llego al final? cuando realice un retorno (return), fuando finaliza todas las acciones y llega al final ( } ), o cuando una Go rutina esta en estado panic. */
package main

import (
	"fmt"
)

type persona struct { //crea un struct
	nombre   string
	apellido string
}

type agenteSecreto struct {
	persona //embebe un struct previamente creado como campo
	lpm     bool
}

func (a agenteSecreto) saludar() { //crea un metodo para el receptor agenteSecreto

	fmt.Println("Buen día", a.nombre, a.apellido) //accede a los campos del receptor usando el id a
}

func main() {

	as := agenteSecreto{
		persona: persona{
			nombre:   "Jessica",
			apellido: "Leonel",
		},
		lpm: true,
	}

	as.saludar() //utiliza el metodo de agenteSecreto llamado saludar
}
