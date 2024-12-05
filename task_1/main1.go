package main

import (
	"fmt"
	"reflect"
)

var numDec int = 42           // Десятичная система
var numOct int = 052          // Восьмеричная система
var numHex int = 0x2A         // Шестнадцатиричная система
var varFlt float64 = 3.14     // Тип float64
var letter string = "Golang"  // Тип string
var complx complex64 = 1 + 2i // Тип complex64
var isState bool = true       // Тип bool

func CheckType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

func main() {
	fmt.Println(CheckType(isState))
}
