package main

import "testing"

func TestSuma(t *testing.T) {
	type prueba struct {
		datos     []int
		respuesta int
	}
	pruebas := []prueba{ //crea un slice de tipo prueba
		prueba{[]int{2, 4, 6}, 12},
		prueba{[]int{1, 5, 2}, 8},
		prueba{[]int{0, -1, 1}, 0},
		prueba{[]int{-10, 4, 20}, 14},
	}

	for _, x := range pruebas {
		v := suma(x.datos...)

		if v != x.respuesta { // de acuerdo al resultado que espero valido
			t.Error("Expected", x.respuesta, "Got", v)
		}
	}

}
