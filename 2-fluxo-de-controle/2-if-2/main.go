package main

import "fmt"

func main() {
	var ehCarro = true
	var valorDoAutomovel = 1000.00

	if ehCarro {
		valorDoAutomovel += 55.55
	}

	fmt.Printf("Valor: %v", valorDoAutomovel)
}
