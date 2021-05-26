package main

import (
	"fmt"
)

type ser interface {
	saludar()
}

type individuo struct {
	nombre string
	genero string
	edad   int
}

type humano struct {
	individuo
	etnia string
}

type animal struct {
	individuo
	raza string
}

func main() {
	perro := animal{
		individuo: individuo{
			nombre: "amy",
			genero: "hembra",
			edad:   4,
		},
		raza: "canino",
	}

	cholo := humano{
		individuo: individuo{
			nombre: "enry",
			genero: "masculino",
			edad:   25,
		},
		etnia: "indigena",
	}

	saludoSer(perro)
	saludoSer(cholo)
}

func (a animal) saludar() {
	fmt.Println("¡Hola!, perrito", a.nombre)
}

func (b humano) saludar() {
	fmt.Println("Buen día, Sr(a).", b.nombre)
}

func saludoSer(s ser) {
	fmt.Println("usando la interface:", s)
}
