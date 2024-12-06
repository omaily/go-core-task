package main

import (
	"bytes"
	"crypto/sha256"
	"encoding/base64"
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

func saltLine(str []rune) []byte {
	var buffer bytes.Buffer
	for i, v := range str {
		if i == len(str)/2 {
			buffer.WriteString("go-2024")
		}
		buffer.WriteRune(v)
	}
	return buffer.Bytes()
}

func main() {
	// 1.1
	arr := make([]interface{}, 0)
	arr = append(arr, 42)
	arr = append(arr, _oct(052))
	arr = append(arr, _hex(0x2A))
	arr = append(arr, 3.14)
	arr = append(arr, 1+2i)
	arr = append(arr, "Golangs")
	arr = append(arr, true)

	// 1.2
	fmt.Print("par.2: ")
	for _, v := range arr {
		fmt.Print(checkType(v), " ")
	}
	fmt.Println()

	//1.3
	fmt.Print("par.3: ")
	str := convertString(arr)
	fmt.Print(str)
	fmt.Println()

	//1.4
	fmt.Print("par.4: ")
	textUTF := []rune(str)
	for _, v := range textUTF {
		fmt.Printf("%c", v)
	}
	fmt.Println()

	//1.5
	fmt.Print("par.5: ")
	textByte := saltLine(textUTF)

	hasher := sha256.New()
	hasher.Write(textByte)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	fmt.Println(sha)
}
