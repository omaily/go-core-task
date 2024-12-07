package task_1

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

func CheckType(v interface{}) string {
	return reflect.TypeOf(v).Name()
}

func ConvertString(arr []interface{}) string {
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
	res := buffer.String()
	return strings.TrimSpace(res)
}

func SaltLine(str []rune) []byte {
	var buffer bytes.Buffer
	for i, v := range str {
		if i == len(str)/2 {
			buffer.WriteString("go-2024")
		}
		buffer.WriteRune(v)
	}
	return buffer.Bytes()
}

func Main() {
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
		fmt.Print(CheckType(v), " ")
	}
	fmt.Println()

	//1.3
	fmt.Print("par.3: ")
	str := ConvertString(arr)
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
	textByte := SaltLine(textUTF)

	hasher := sha256.New()
	hasher.Write(textByte)
	sha := base64.URLEncoding.EncodeToString(hasher.Sum(nil))
	fmt.Println(sha)
}
