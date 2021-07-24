package decir

import (
	"fmt"
	"testing"
)

func TextSaludar(t *testing.T) {
	s := Saludar("Jessica")

	if s != "Bienvenido querid@ Jessica" {
		t.Error("Expected", "Bienvenido querid@ Jessica", "Got", s)
	}
}

func ExampleSaludar() {
	fmt.Println(Saludar("Jessica"))
	//Output:
	//Bienvenido querid@ Jessica
}

func BenchmarkSaludar(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Saludar("Jessica")
	}
}
