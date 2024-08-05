package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	var int8 uint8 = 127
	var inteiro int
	inteiro = int(int8)
	println(inteiro)

	var f32 float32 = 1.23
	inteiro = int(f32)
	println(inteiro)

	b, _ := strconv.ParseBool("true")
	fmt.Println(reflect.TypeOf(b), b)
	fmt.Printf("O tipo da variável é %T", b)
}
