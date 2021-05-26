package main

import (
	"fmt"
)

func main() {
	s := struct {
		nombre  string
		amigos  map[string]int //campo de tipo mapa
		bebidas []string       //campo de tipo slice
	}{
		nombre: "Jorge",
		amigos: map[string]int{
			"andrea": 30,
			"luisa":  31,
		},
		bebidas: []string{"vino tinto", "cerveza artesanal", "aguardiente"},
	}

	fmt.Println(s.nombre)
	fmt.Println(s.amigos)
	fmt.Println(s.bebidas)

	fmt.Println("-----------------")

	for i, v := range s.amigos {
		fmt.Println(i, v)
	}

	fmt.Println("-----------------")

	for _, value := range s.bebidas {
		fmt.Println(value)
	}
}
