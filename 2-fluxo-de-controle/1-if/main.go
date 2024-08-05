package main

import "fmt"

func main() {
	salario := 900.00
	var salarioComBonus float64

	salarioComBonus = salario

	if salario < 1000.00 {
		salarioComBonus = salario + 200.00
	}

	fmt.Printf("Salário: %v", salarioComBonus)
}
