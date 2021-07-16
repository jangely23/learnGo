/*
podemos especifiar si el canal solo envia o solo recibe
se muestra en el codigo comentado las situaciones que romperian le programa
esta ruptura la genera un mal uso del canal.
*/

package main

import "fmt"

func main() {

	//ejemplos no funcionales de tipos de canales

	ca2 := make(<-chan int) //crea un canal recibe only
	ca3 := make(chan<- int) //crea un tipo send only
	/*
		//ERROR ca2 solo envia, no puede recibir
		ca2 <- 92 //envia informacion a el canal.
	*/

	go func() { //se requiere receptor debido a que no tiene bufer
		ca3 <- 92
	}()

	go func() { //se requiere receptor debido a que no tiene buffer
		fmt.Println(<-ca2)
	}()

	/*
		//ERROR ca3 solo recibe, no puede enviar
		fmt.Println(<-ca3)
		fmt.Println(<-ca3) //obtiene la informacion del canal
	*/
	fmt.Println("-------------------")
	fmt.Printf("ca2 %T\n", ca2)
	fmt.Printf("ca3 %T\n", ca3)

}
