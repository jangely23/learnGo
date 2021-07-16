/*agarra un trabajo secuencial, dividirlo en porciones y que una gorutina tabaje en una porcion de ese trabajo*/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	c1 := make(chan int)
	c2 := make(chan int)

	go agregar(c1)

	go fanOutIn(c1, c2)

	for v := range c2 { //recibe los valores
		fmt.Println(v)
	}

	fmt.Println("Finalizando.")
}

func agregar(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}

func fanOutIn(c1, c2 chan int) {
	var wg sync.WaitGroup
	const gorutinas = 10
	wg.Add(gorutinas)

	for i := 0; i < gorutinas; i++ { //lanza una go rutina por cada porcion
		go func() {
			for v := range c1 {
				func(v2 int) {
					c2 <- trabajoConsumeTiempo(v2) //ejecuta un trabajo
				}(v)
			}
			wg.Done()
		}()
	}
	wg.Wait()
	close(c2)
}

func trabajoConsumeTiempo(n int) int {
	time.Sleep(time.Microsecond * time.Duration(rand.Intn(500)))
	return n + rand.Intn(1000)
}
