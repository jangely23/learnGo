//canal si bufer sin bloqueo

package main

import "fmt"

func main() {
	//unbuffered channels (sin búfer) solo se utilizan con una gorutina que transmite y otra recibe
	ca := make(chan int) //crea un canal

	go func() {
		ca <- 42 //la go rutina de la funcion anonima es el receptor.
	}()

	fmt.Println(<-ca) //la gorutina principal (main) recibe a ca
}
