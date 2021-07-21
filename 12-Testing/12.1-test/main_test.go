package main

import "testing"

func TestSuma(t *testing.T) {
	v := suma(20, 8)

	if v != 28 { // de acuerdo al resultado que espero valido
		t.Error("Expected", 28, "Got", v)
	}
}
