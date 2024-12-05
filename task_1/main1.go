package main

import (
	"fmt"
	"reflect"
	"strings"
)

type _oct int
type _hex int

func checkType(v interface{}) string {
	return reflect.TypeOf(v).Name()
}

func convertString(arr []interface{}) string {
	var buffer strings.Builder
	for _, v := range arr {
		switch reflect.TypeOf(v).Name() {
		case "_oct":
			buffer.WriteString(fmt.Sprintf("O%o ", v))
		case "_hex":
			buffer.WriteString(fmt.Sprintf("%#x ", v))
		default:
			buffer.WriteString(fmt.Sprint(v, " "))
		}
	}
	return buffer.String()
}

func main() {
	// 1.1
	var oct _oct = 052
	var hex _hex = 0x2A
	arr := make([]interface{}, 0)
	arr = append(arr, 42)
	arr = append(arr, oct)
	arr = append(arr, hex)
	arr = append(arr, 3.14)
	arr = append(arr, 1+2i)
	arr = append(arr, "Golang")
	arr = append(arr, true)

	// 1.2
	for _, v := range arr {
		fmt.Println(checkType(v))
	}

	//1.3
	fmt.Println(convertString(arr))

}
