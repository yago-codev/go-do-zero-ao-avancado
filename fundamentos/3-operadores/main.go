package main

import (
	"fmt"
	"reflect"
)

func main() {
	var1 := 10
	var2 := 20
	total := var1 + var2
	fmt.Println(reflect.TypeOf(total))

	texto1 := "primeira palavra "
	texto2 := "segunda palavra"
	unirTextos := texto1 + texto2
	fmt.Println(unirTextos)
}
