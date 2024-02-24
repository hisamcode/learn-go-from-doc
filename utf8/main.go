package main

import (
	"fmt"
	"unicode/utf8"
)

func triangle() {
	for i := 0; i < 5; i++ {
		a := ""
		for z := 0; z < i; z++ {
			a += "*"
		}
		fmt.Println(a)
	}
}

func main() {
	triangle()
	return
	test := "ಚa" // is 3 bytes
	fmt.Println(test, " len", len(test))
	// output len is 3
	r := []rune(test)
	fmt.Println([]byte(test))
	fmt.Println(utf8.RuneCount([]byte(test)))
	fmt.Println("rune", r, len(r))

	for i := 0; i < len(test); i++ {
		fmt.Println(test[i])
	}
}
