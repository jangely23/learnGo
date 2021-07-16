//canal si bufer y sin receptor bloquea la gorutina principal

package main

import "fmt"

func main() {
	//unbuffered channels (sin búfer) solo se utilizan con una gorutina que transmite y otra recibe
	ca := make(chan int) //crea un canal

	ca <- 42 //envia 42 al canal ca como no tiene quien reciba se bloquea el codigo.

	fmt.Println(<-ca) //la gorutina principal (main) recibe a ca
}
