package main

import (
	"context"
	"fmt"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancela cuando hemos finalizado el main

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
}

func gen(ctx context.Context) <-chan int { //recibe el contexto y retorna un canal
	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return // retornando para que no se fuge la gorutina (finaliza)
			case dst <- n:
				n++
			}
		}
	}()
	return dst
}
