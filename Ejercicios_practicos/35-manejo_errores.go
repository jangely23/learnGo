package main

import "fmt"

type errorPer struct {
	mensaje string
}

func (ep errorPer) Error() string { //implementa la interfaz builtin.error
	return fmt.Sprintf("aqui esta el error: %v", ep.mensaje)
}

func main() {
	ep2 := errorPer{
		mensaje: "Requiero vacaciones",
	}

	foo(ep2)
}

func foo(e error) {
	fmt.Println("foo corrio ", e, "\n", e.(errorPer).mensaje) //asesion de tipo
}
