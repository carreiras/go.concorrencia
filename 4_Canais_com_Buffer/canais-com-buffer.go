package main

import (
	"fmt"
)

func main() {
	canal := make(chan string, 2)
	canal <- "Olá Mundo!"
	canal <- "Programando do Go!"

	mensagem1 := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem1)
	fmt.Println(mensagem2)
}
