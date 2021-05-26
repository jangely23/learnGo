package main

import (
	"fmt"
)

func main() {

	p1 := struct { //struct anonimo
		nombre string
		edad   int
		genero string
	}{
		nombre: "jessica", //valores campos struct anonimo
		edad:   31,
		genero: "femenino",
	}

	fmt.Println(p1)
}
