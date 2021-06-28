/*
Ordena de acuerdo a lo que se le indique ejemlo: nombre, edad, etc.
*/

package main

import (
	"fmt"
	"sort"
)

type Persona struct {
	Nombre string
	Edad   int
}

//permite orderar personalizado en este caso por edad

type PorEdad []Persona

func (a PorEdad) Len() int           { return len(a) }
func (a PorEdad) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a PorEdad) Less(i, j int) bool { return a[i].Edad < a[j].Edad } //compara cual es menor que cual
func main() {
	p1 := Persona{"Jessica", 31}

	p2 := Persona{"Zully", 29}

	p3 := Persona{"Lauren", 14}

	p4 := Persona{"Danelly", 30}

	p5 := Persona{"Martha", 56}

	personas := []Persona{p1, p2, p3, p4, p5}

	fmt.Println(personas)

	sort.Sort(PorEdad(personas)) //tiene adjunta la interfaz con los metodos Len, Swap, Less.
	fmt.Println(personas)
}
