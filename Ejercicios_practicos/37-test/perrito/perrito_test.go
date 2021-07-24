package perrito

import (
	"fmt"
	"testing"
)

func TestEdad(t *testing.T) {
	e := Edad(10)

	if e != 70 {
		t.Error("Edad invalida", "Expected", 70, "Got", e)
	}
}

func ExampleEdad() {
	fmt.Println(Edad(10))
	//Output:
	//70
}

func BenchmarkEdad(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Edad(10)
	}
}

//Funcion EdadDos
func TestEdadDos(t *testing.T) {
	e := EdadDos(10)

	if e != 70 {
		t.Error("Edad invalida", "Expected", 70, "Got", e)
	}
}

func ExampleEdadDos() {
	fmt.Println(EdadDos(10))
	//Output:
	//70
}

func BenchmarkEdadDos(b *testing.B) {
	for i := 0; i < b.N; i++ {
		EdadDos(10)
	}
}
