package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func checkType(v interface{}) reflect.Type {
	return reflect.TypeOf(v)
}

func convertString(arr []interface{}) string {
	var buffer bytes.Buffer
	for _, v := range arr {
		buffer.WriteString(fmt.Sprint(v, " "))
	}
	return buffer.String()
}

func main() {
	// 1.1
	arr := make([]interface{}, 0)
	arr = append(arr, 42)
	arr = append(arr, 052)
	arr = append(arr, 0x2A)
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
