package main

import "fmt"

func main() {
	par := make(chan int)
	impar := make(chan int)
	salir := make(chan bool) //se cambia el tipo de dato a booleano

	//enviar
	go enviar(par, impar, salir)

	//recibir
	recibir(par, impar, salir)

	fmt.Println("Finalizado...")

}

func enviar(par, impar chan<- int, salir chan<- bool) {
	for j := 0; j < 100; j++ {
		if j%2 == 0 {
			par <- j
		} else {
			impar <- j
		}
	}
	close(salir)
}

func recibir(par, impar <-chan int, salir chan<- bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("desde el canal Par: ", v)
		case v := <-impar:
			fmt.Println("desde el canal impar: ", v)
		case i, ok := <-salir: //idioma, coma ok
			if !ok {
				fmt.Println("desde el coma ok: ", i)
				return
			} else {
				fmt.Println("desde el coma ok: ", i)
			}

		}
	}
}
