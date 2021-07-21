package main

import (
	"fmt"

	"github.com/learnGo/11-Documentacion/11.2-paquete_perro/perro"
)

type canino struct {
	nombre string
	edad   int
}

func main() {
	c1 := canino{
		nombre: "amy",
		edad:   perro.Edad(6),
	}

	fmt.Println("la edad humana es: ", c1.edad)
}
